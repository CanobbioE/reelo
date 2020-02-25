package interactor

import (
	"context"
	"fmt"
	"strings"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
	"github.com/CanobbioE/reelo/backend/utils/solvers"
	"golang.org/x/sync/errgroup"
)

// ListNamesakes recognize and returns the majority of namesakes. A namesake
// is a player that has two or more results that are impossible for a single person
// to achive. E.g. if it partecipates in two different categories the same year.
func (i *Interactor) ListNamesakes(page, size int) ([]usecases.Namesake, utils.Error) {
	var namesakes []usecases.Namesake
	errs, ctx := errgroup.WithContext(context.Background())
	players, err := i.PlayerRepository.FindAll(ctx, page, size)
	if err != nil {
		return namesakes, utils.NewError(err, "E_DB_FIND", 500)
	}
	for _, player := range players {
		player := player
		errs.Go(func() error {
			history, years, err := i.HistoryRepository.FindByPlayerIDOrderByYear(ctx, player.ID)
			if err != nil {
				return err
			}
			comment := domain.Comment{Text: ""}
			if i.CommentRepository.CheckExistenceByPlayerID(context.Background(), player.ID) {
				i.Logger.Log("ListNamesakes: cannot check comment existence: %v", err)
				comment, err = i.CommentRepository.FindByPlayerID(context.Background(), player.ID)
				if err != nil {
					i.Logger.Log("ListNamesakes: cannot find comment for player %v: %v", player.ID, err)
					return err
				}
			}

			// solvedPlayer is an array of dto.Namesake objects
			solvedPlayer, err := solvePlayer(player, history, years, comment, i.UpdateNamesake)
			if solvedPlayer != nil {
				namesakes = append(namesakes, solvedPlayer...)
			}
			return err
		})
	}

	// if errs.Wait is nil then the returned error will be nil
	return namesakes, utils.NewError(errs.Wait(), "E_GENERIC", 500)
}

// SolvePlayer executes the logic to determinate if a player has namesakes
func solvePlayer(player domain.Player, history usecases.HistoryByYear, years []int,
	comment domain.Comment, autoSolver func(n usecases.Namesake) utils.Error) ([]usecases.Namesake, error) {
	ss := solvers.New()
	var namesakes []usecases.Namesake
	namesakes = nil

	// years is sorted
	for _, y := range years {
		for _, result := range history[y] {
			for ss.Next() {
				if ss.Current().CanAccept(solvers.SlimPartecipation(result)) {
					ss.AppendToCurrent(solvers.SlimPartecipation(result))
					break
				} else if !ss.HasNext() {
					ss.NewSolver(solvers.SlimPartecipation(result))
				}
			}
			ss.ResetCursor()
		}
	}

	if ss.Size() > 1 {
		ss.ResetCursor()
		for i := 0; ss.Next(); i++ {
			solver := ss.Current()
			namesakes = append(namesakes, usecases.Namesake{
				Solver:  *solver,
				ID:      i,
				Comment: comment,
				Player:  player,
			})
		}
		if ss.ShouldBeManual() {
			return namesakes, nil
		}
		for _, n := range namesakes {
			err := autoSolver(n)
			if !err.IsNil {
				return namesakes, fmt.Errorf(err.Message)
			}
			return []usecases.Namesake{}, nil
		}
	}
	return namesakes, nil
}

// UpdateNamesake changes a player history by assigning the history specified by the given
// Namesake to a newly created player. The new player has the same name/surname of the old player
// (hence it's a namesake) but has a different accent
func (i *Interactor) UpdateNamesake(n usecases.Namesake) utils.Error {
	ctx := context.Background()
	oldID := n.Player.ID
	oldHistories, years, err := i.HistoryRepository.FindByPlayerIDOrderByYear(ctx, oldID)
	if err != nil {
		i.Logger.Log("UpdateNamesake: cannot find history for player %v: %v", oldID, err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}
	if len(years) == 0 || len(oldHistories) == 0 {
		return utils.NewNilError()
	}

	var accentID int
repeat:
	accent := fmt.Sprintf("%d %s %d", years[0], n.Solver[0].City, accentID)
	p := n.Player
	p.Accent = accent
	newID, err := i.PlayerRepository.Store(ctx, p)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			accentID++
			goto repeat
		}
		i.Logger.Log("UpdateNamesake: cannot store player %v: %v", p.ID, err)
		return utils.NewError(err, "E_DB_STORE", 500)

	}

	// Edit the old player's history by reassigning her/his results to the new player's ID
	for _, newEntry := range n.Solver.ToHistory() {
		if oldHistory, ok := oldHistories[newEntry.Year]; ok {
			for _, oldEntry := range oldHistory {
				if oldEntry.IsEqual(usecases.SlimPartecipation(newEntry)) {
					oldEntryID, err := i.ResultRepository.FindIDByPlayerIDAndGameYearAndCategory(ctx, oldID, oldEntry.Year, oldEntry.Category)
					if err != nil {
						i.Logger.Log("UpdateNamesake: cannot find result's ID: %v", err)
						return utils.NewError(err, "E_DB_FIND", 500)
					}
					if err := i.PartecipationRepository.UpdatePlayerIDByGameID(ctx, int(newID), oldEntryID); err != nil {
						i.Logger.Log("UpdateNamesake: cannot update partecipation player's ID: %v", err)
						return utils.NewError(err, "E_DB_UPDATE", 500)
					}
				}
			}

		}
	}
	_, err = i.PartecipationRepository.FindByPlayerID(ctx, oldID)
	if err != nil {
		if err.Error() != "no values in result set" {
			i.Logger.Log("UpdateNamesake: cannot find partecipation: %v", err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}
		if err := i.PlayerRepository.DeleteByID(ctx, oldID); err != nil {
			i.Logger.Log("UpdateNamesake: cannot delete player %v: %v", oldID, err)
			return utils.NewError(err, "E_DB_DELETE", 500)
		}
	}
	return utils.NewNilError()
}
