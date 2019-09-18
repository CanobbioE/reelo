package services

import (
	"context"
	"log"

	rdb "github.com/CanobbioE/reelo/backend/db"
	solvers "github.com/CanobbioE/reelo/backend/utils/solvers"
)

// PurgeNamesakes tries to identify and split players that are namesakes.
func PurgeNamesakes() error {
	count := 0
	db := rdb.NewDB()
	defer db.Close()

	log.Println("Started purging...")

	players, err := db.AllPlayers(context.Background())
	if err != nil {
		return err
	}
	for _, player := range players {
		ss := solvers.New()

		// History is a map[year]results where results is an array of results
		history, years, err := db.AnalysisHistory(context.Background(), player.Name, player.Surname)
		if err != nil {
			return err
		}

		// years is sorted
		for _, y := range years {
			for _, result := range history[y] {
				for ss.Next() {
					if ss.Current().CanAccept(result) {
						ss.AppendToCurrent(result)
					} else if !ss.HasNext() {
						ss.NewSolver(result)
					}
				}
				ss.ResetCursor()
			}
		}
		if ss.Size() > 1 {
			count++
			log.Printf("Found namesake: %v %v: \n%v\n", player.Name, player.Surname, ss)
		}
	}

	log.Printf("Found %v namesakes", count)

	return nil
}
