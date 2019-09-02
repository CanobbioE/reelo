package services

import (
	"context"

	rdb "github.com/CanobbioE/reelo/backend/db"
)

// PurgeNamesakes tries to identify and split players that are namesakes.
func PurgeNamesakes() error {
	db := rdb.NewDB()
	defer db.Close()

	players, err := db.AllPlayers(context.Background())
	if err != nil {
		return err
	}
	for _, player := range players {
		history, err := db.AnalysisHistory(context.Background(), player.Name, player.Surname)
		if err != nil {
			return err
		}

		var solvers [][]rdb.History
		for _, vv := range history {
			for i, v := range vv {
				if i == 0 {
					solvers[0] = []rdb.History{v}
				} else {
					for i, solver := range solvers {
						if isOk(solver, v) {
							solvers[i] = append(solvers[i], v)
						} else if i == len(solvers)-1 {
							solvers = append(solvers, []rdb.History{v})
						}
					}
				}
			}
		}

	}

	return nil
}

func isOk(solver []rdb.History, current rdb.History) bool {
	/*
		check 2 results in one year,
		check category growth

		last := solver[len(solver)-1]
			if (last.Year == current.Year) &&
				(last.IsParis || current.IsParis) &&
				(last.Category == current.Category) {
				return true
			} else {
				if last.Year !=  current.Year &&
				elo.CategoryFromString(last.Category) = elo.CategoryFromString(current.Category) {

				}

			}
			// years do not corresponds
			// if they do, one must be international
			// category corresponds to year
			// places corresponds?
			return false
	*/
	return false

}

/* CATEGORY GROWTH:
ce: 0
cm: 1,2
C1: 3,4
C2: 4,5
L1: 6,7,8
L2: 9, 10, 11
GP: 12,...,99
*/
