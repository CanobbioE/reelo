package interactor

import (
	"context"
	"fmt"
	"strings"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils/solvers"
	"golang.org/x/sync/errgroup"
)

// ListNamesakes recognize and returns the majority of namesakes. A namesake
// is a player that has two or more results that are impossible for a single person
// to achive. E.g. if it partecipates in two different categories the same year.
func (i *Interactor) ListNamesakes(page, size int) ([]usecases.Namesake, error) {
	var namesakes []usecases.Namesake
	errs, ctx := errgroup.WithContext(context.Background())
	players, err := i.PlayerRepository.FindAll(ctx, page, size)
	if err != nil {
		return namesakes, err
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
				comment, err = i.CommentRepository.FindByPlayerID(context.Background(), player.ID)
				if err != nil {
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

	return namesakes, errs.Wait()
}

// SolvePlayer executes the logic to determinate if a player has namesakes
func solvePlayer(player domain.Player, history usecases.HistoryByYear, years []int,
	comment domain.Comment, autoSolver func(n usecases.Namesake) error) ([]usecases.Namesake, error) {
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
			if err != nil {
				return namesakes, err
			}
			return []usecases.Namesake{}, nil
		}
	}
	return namesakes, nil
}

// UpdateNamesake changes a player history by assigning the history specified by the given
// Namesake to a newly created player. The new player has the same name/surname of the old player
// (hence it's a namesake) but has a different accent
func (i *Interactor) UpdateNamesake(n usecases.Namesake) error {
	ctx := context.Background()
	oldID := n.Player.ID
	oldHistories, years, err := i.HistoryRepository.FindByPlayerIDOrderByYear(ctx, oldID)
	if err != nil {
		return err
	}
	if len(years) == 0 || len(oldHistories) == 0 {
		return nil
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
		return err
	}

	// Edit the old player's history by reassigning her/his results to the new player's ID
	for _, newEntry := range n.Solver.ToHistory() {
		if oldHistory, ok := oldHistories[newEntry.Year]; ok {
			for _, oldEntry := range oldHistory {
				if oldEntry.IsEqual(usecases.SlimPartecipation(newEntry)) {
					oldEntryID, err := i.ResultRepository.FindIDByPlayerIDAndGameYearAndCategory(ctx, oldID, oldEntry.Year, oldEntry.Category)
					if err != nil {
						return err
					}
					if err := i.PartecipationRepository.UpdatePlayerIDByGameID(ctx, int(newID), oldEntryID); err != nil {
						return err
					}
				}
			}

		}
	}
	_, err = i.PartecipationRepository.FindByPlayerID(ctx, oldID)
	if err != nil {
		if err.Error() != "no values in result set" {
			return err
		}
		if err := i.PlayerRepository.DeleteByID(ctx, oldID); err != nil {
			return err
		}
	}
	return nil
}
