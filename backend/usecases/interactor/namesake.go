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
// to achieve. E.g. if it participates in two different categories the same year.
func (i *Interactor) ListNamesakes(page, size int) ([]usecases.Namesake, utils.Error) {
	var namesakes []usecases.Namesake
	errs, ctx := errgroup.WithContext(context.Background())
	players, err := i.PlayerRepository.FindAll(ctx, page, size)
	if err != nil {
		return namesakes, utils.NewError(err, "E_DB_FIND", 500)
	}

	for _, player := range players {
		player := player // we need this here
		errs.Go(func() error {
			history, years, err := i.HistoryRepository.FindByPlayerIDOrderByYear(ctx, player.ID)
			if err != nil {
				return err
			}
			var comment domain.Comment
			commentExists := i.CommentRepository.CheckExistenceByPlayerID(context.Background(), player.ID)
			if commentExists {
				comment, err = i.CommentRepository.FindByPlayerID(context.Background(), player.ID)
				if err != nil {
					i.Logger.Log("ListNamesakes: cannot find comment for player %v: %v", player.ID, err)
					return err
				}
			}

			// solvedPlayer is an array of Namesakes
			solvedPlayer, err := solvePlayer(player, history, years, comment, i.UpdateNamesake)
			if err != nil {
				i.Logger.Log("ListNamesakes: cannot solve player %v: %v", player.ID, err)
				return err
			}
			if solvedPlayer != nil {
				namesakes = append(namesakes, solvedPlayer...)
			}
			return nil
		})
	}

	// if errs.Wait is nil then the returned error will be nil
	return namesakes, utils.NewError(errs.Wait(), "E_GENERIC", 500)
}

// SolvePlayer executes the logic to determinate if a player has namesakes
func solvePlayer(
	player domain.Player, history usecases.HistoryByYear,
	years []int, comment domain.Comment,
	autoSolver func(n usecases.Namesake) utils.Error,
) ([]usecases.Namesake, error) {

	ss := solvers.New()
	var namesakes []usecases.Namesake
	namesakes = nil

	// years is sorted
	for _, y := range years {
		for _, result := range history[y] {
			for ss.Next() {
				if ss.Current().CanAccept(solvers.SlimParticipation(result)) {
					ss.AppendToCurrent(solvers.SlimParticipation(result))
					break
				} else if !ss.HasNext() {
					ss.NewSolver(solvers.SlimParticipation(result))
				}
			}
			ss.ResetCursor()
		}
	}

	// a player that has multiple possible histories
	// is what we want to return
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
			return nil, nil
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
		i.Logger.Log("UpdateNamesake: cannot store player %v: %v", p.Accent, err)
		return utils.NewError(err, "E_DB_STORE", 500)
	}

	// Edit the old player's history by reassigning her/his results to the new player's ID
	// according to what the solver specifies
	for _, newEntry := range n.Solver.ToHistory() {
		if oldHistory, ok := oldHistories[newEntry.Year]; ok {
			for _, oldEntry := range oldHistory {
				if oldEntry.IsEqual(usecases.SlimParticipation(newEntry)) {
					oldEntryID, err := i.ResultRepository.FindIDByPlayerIDAndGameYearAndCategory(ctx, oldID, oldEntry.Year, oldEntry.Category)
					if err != nil {
						i.Logger.Log("UpdateNamesake: cannot find result's ID: %v", err)
						return utils.NewError(err, "E_DB_FIND", 500)
					}
					if err := i.ParticipationRepository.UpdatePlayerIDByResultID(ctx, int(newID), oldEntryID); err != nil {
						i.Logger.Log("UpdateNamesake: cannot update participation player's ID: %v", err)
						return utils.NewError(err, "E_DB_UPDATE", 500)
					}
				}
			}

		}
	}

	// Remove the old player entry if it has no entry in its history
	var mustDelete bool
	history, err := i.ParticipationRepository.FindByPlayerID(ctx, oldID)
	if err != nil {
		if strings.Contains(err.Error(), "no values in result set") {
			mustDelete = true
			goto delete
		}
		i.Logger.Log("UpdateNamesake: cannot find participation: %v", err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}

delete:
	if len(history) == 0 || mustDelete {
		if err := i.PlayerRepository.DeleteByID(ctx, oldID); err != nil {
			i.Logger.Log("UpdateNamesake: cannot delete player %v: %v", oldID, err)
			return utils.NewError(err, "E_DB_DELETE", 500)
		}
	}

	return utils.NewNilError()
}
