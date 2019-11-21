package services

import (
	"context"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/dto"
	solvers "github.com/CanobbioE/reelo/backend/utils/solvers"
	"golang.org/x/sync/errgroup"
)

// GetNamesakes recognize and returns the majority of namesakes. A namesake
// is a player that has two or more results that are impossible for a single person
// to achive. E.g. if it partecipates in two different categories the same year.
func GetNamesakes(page, size int) ([]dto.Namesake, error) {
	db := rdb.Instance()
	var namesakes []dto.Namesake
	errs, ctx := errgroup.WithContext(context.Background())
	players, err := db.AllPlayers(ctx, page, size)
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
func solvePlayer(player rdb.Player) ([]dto.Namesake, error) {
	ss := solvers.New()
	ctx := context.Background()
	var namesakes []dto.Namesake
	namesakes = nil

	// History is a map[year]results where results is an array of results
	db := rdb.Instance()
	history, years, err := db.AnalysisHistory(ctx, player.Name, player.Surname)
	if err != nil {
		return namesakes, err
	}
	playerID, err := db.PlayerID(ctx, player.Name, player.Surname)
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
			namesakes = append(namesakes, dto.Namesake{
				Name:     player.Name,
				Surname:  player.Surname,
				PlayerID: playerID,
				Solver:   *solver,
				ID:       i,
			})
		}
	}
	return namesakes, nil
}
