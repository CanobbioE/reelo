package interactor

import (
	"context"

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
			// solvedPlayer is an array of dto.Namesake objects
			solvedPlayer, err := solvePlayer(player)
			if solvedPlayer != nil {
				namesakes = append(namesakes, solvedPlayer...)
			}
			return err
		})
	}

	return namesakes, errs.Wait()
}

// SolvePlayer executes the logic to determinate if a player has namesakes
func solvePlayer(player domain.Player) ([]usecases.Namesake, error) {
	ss := solvers.New()
	ctx := context.Background()
	var namesakes []usecases.Namesake
	namesakes = nil

	history, years, err := i.HistoryRepository.FindByPlayerID(ctx, player.ID)
	if err != nil {
		return namesakes, err
	}
	comment, err := i.CommentRepository.CheckExistenceByPlayerID(context.Background(), player.ID)
	if err != nil {
		return namesakes, err
	}

	// years is sorted
	for _, y := range years {
		for _, result := range history[y] {
			for ss.Next() {
				if ss.Current().CanAccept(result) {
					ss.AppendToCurrent(result)
					break
				} else if !ss.HasNext() {
					ss.NewSolver(result)
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
			})
		}
	}
	return namesakes, nil
}

// UpdateNamesake changes a player history
func (i *Interactor) UpdateNamesake(n usecases.Namesake) error {
	ctx := context.Background()
	oldID := n.Player.ID
	// accent := rdb.CreateAccent(history[0].Year, accentID, history[0].City)
	newID, err := i.PlayerRepository.Store(ctx, namesake.Player)
	if err != nil {
		return err
	}

	// Edit the old player's history by reassigning her/his results to the new player's ID
	/*
		err = db.HistorySwitcheroo(ctx, oldID, newID, history)
		if err != nil {
			return err
		}
	*/
	return nil
}
