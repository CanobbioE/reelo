package elo

import (
	"math"
	"time"

	rdb "github.com/CanobbioE/reelo/db"
)

var db = rdb.NewDB()

const STARTING_YEAR = 2002
const K_EXERCISES = 20
const P_FINAL = 1.5
const MULTIPLICATIVE_FACTOR = 10000
const ANTI_EXPLOIT = 0.9
const NP_PENALTY = 0.9

// Reelo returns the points for a given user calculated with a custom algorithm.
// TODO: Reelo should not make DB calls
func Reelo(name, surname string) (reelo float64) {
	var count, lastKnownYear int
	var lastKnownCategory string

	// TODO: fix the for loop's range
	for _, year := range []int{2002, 2018} {
		// Variables names are chosen accordingly to the formula provided by the scientific committee
		category := db.GetCategory(name, surname, year)

		t := StartOfCategory(year, category)
		n := EndOfCategory(year, category)
		eMax := t - n + 1
		dMax := MaxScoreForCategory(year, category)
		d := db.GetScore(name, surname, year)
		e := db.GetExercises(name, surname, year)

		baseScore := float64(d + e*K_EXERCISES)

		// if FINALE INTERNAZIONALE {
		//    baseScore = baseScore * P_FINAL
		// }

		for i := 1; i <= t-1; i++ {
			errorFactor := float64(1-d) + float64(e*K_EXERCISES)/float64((K_EXERCISES*eMax+dMax))
			difficultyFactor := float64(i) / float64(n+1)
			nonResolutionProbability := 1 - errorFactor*difficultyFactor
			baseScore += float64(K_EXERCISES+i) * nonResolutionProbability
		}

		baseScore = baseScore / db.GetAvgScoresOfCategories(year)
		baseScore = baseScore * MULTIPLICATIVE_FACTOR

		// TODO: actually check for category promotion
		if lastKnownCategory > category {
			oldAvg := db.GetAvgScore(year, category)
			newAvg := db.GetAvgScore(year, lastKnownCategory)
			newMax := db.GetMaxScore(year, lastKnownCategory)
			scoreOfYear := baseScore

			convertedScore := scoreOfYear * newAvg / oldAvg
			if convertedScore > newMax {
				convertedScore = newMax
			}

			baseScore = convertedScore
		}

		agingFactor := 1 - 5/72*math.Log2(float64(lastKnownYear-year+1))
		baseScore = baseScore * agingFactor

		reelo = reelo + baseScore
		count++
	}

	reelo = reelo / float64(count)

	if count == 1 && lastKnownYear == time.Now().Year() {
		reelo = reelo * ANTI_EXPLOIT
	}

	// if NO RESULT IN THE LAST YEAR {
	//    reelo = reelo * NP_PENALTY
	// }
	return reelo
}
