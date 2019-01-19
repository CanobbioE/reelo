package elo

import (
	"context"

	rdb "github.com/CanobbioE/reelo/db"
)

var db = rdb.NewDB()

const STARTING_YEAR = 2002

// Reelo returns the points for a given user calculated with a custom algorithm.
func Reelo(name, surname string) (reelo int) {
	defer db.Close()
	results := db.GetResults(context.Background(), name, surname)
	lastYear := STARTING_YEAR
	for _, r := range results {
		// if y isEmpty
		if r.Year-lastYear > 1 {
			// reelo = avg(pastFullYears)/2
		} else { // if y isFull
			// cc = sum(Dmax)/Emax * Tmax
			// reelo = (cc * sum(D+E*K)/avgOfAvgCat) * 10000
			// 	if y is first full
			// 	reelo = (reelo + reelo/2)/2
		}
		// K_aging = 1 - 5/72 * logBase2(anno-annoi+1)
		// reelo = reelo * Kaging
	}
	return reelo
}
