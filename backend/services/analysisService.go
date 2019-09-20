package services

import (
	"context"
	"log"

	rdb "github.com/CanobbioE/reelo/backend/db"
	solvers "github.com/CanobbioE/reelo/backend/utils/solvers"
	"golang.org/x/sync/errgroup"
)

// PurgeNamesakes tries to identify and split players that are namesakes.
func PurgeNamesakes() error {
	db := rdb.Instance()
	errs, ctx := errgroup.WithContext(context.Background())
	players, err := db.AllPlayers(ctx)
	if err != nil {
		return err
	}

	log.Println("Started purging...")
	for _, player := range players {
		player := player
		errs.Go(func() error {
			return solvePlayer(player)
		})
	}

	return errs.Wait()
}

// SolvePlayer executes the logic to determinate if a player has namesakes
func solvePlayer(player rdb.Player) error {
	ss := solvers.New()
	ctx := context.Background()

	// History is a map[year]results where results is an array of results
	db := rdb.Instance()
	history, years, err := db.AnalysisHistory(ctx, player.Name, player.Surname)
	if err != nil {
		return err
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
		log.Printf("Found namesake: %v %v: \n%v\n", player.Name, player.Surname, ss)
	}
	return nil
}
