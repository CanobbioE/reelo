package elo

import (
	"math"
	"time"

	rdb "github.com/CanobbioE/reelo/backend/db"
)

var db = rdb.NewDB()

// TODO: fetch this from db
const STARTING_YEAR = 2002
const K_EXERCISES = 20
const P_FINAL = 1.5
const MULTIPLICATIVE_FACTOR = 10000
const ANTI_EXPLOIT = 0.9
const NP_PENALTY = 0.9

// Reelo returns the points for a given user calculated with a custom algorithm.
// TODO: Handle Internationa Finals
// TODO: from the DB recover: lastKnownYear, lastKnownCategory
func Reelo(name, surname string) (reelo float64) {
	var count, lastKnownYear int
	var lastKnownCategory string
	var intenationalFinal bool

	// TODO: fix the for loop's range
	for _, year := range []int{2002, 2018} {
		// Variables names are chosen accordingly to the formula
		// provided by the scientific committee
		category := db.GetCategory(name, surname, year)

		t := StartOfCategory(year, category)
		n := EndOfCategory(year, category)
		eMax := t - n + 1
		dMax := MaxScoreForCategory(year, category)
		d := db.GetScore(name, surname, year)
		e := db.GetExercises(name, surname, year)

		baseScore := float64(d + e*K_EXERCISES)

		if intenationalFinal {
			baseScore = baseScore * P_FINAL
		}

		// Categories homogenization:
		// For each exercises a player is not supposed to solve we calculate
		// her/his probabilty of solving it.
		for i := 1; i <= t-1; i++ {
			errorFactor := 1 - float64(d+e*K_EXERCISES)/float64((K_EXERCISES*eMax+dMax))
			difficultyFactor := float64(i) / float64(n+1)
			nonResolutionProbability := 1 - errorFactor*difficultyFactor
			baseScore += float64(K_EXERCISES+i) * nonResolutionProbability
		}

		baseScore = baseScore / db.GetAvgScoresOfCategories(year)
		baseScore = baseScore * MULTIPLICATIVE_FACTOR

		// TODO: actually check for category promotion
		// Category promotion:
		// If this year the player's category is inferior to the category she/he played
		// most recently then we convert this year score to the most recent category
		if lastKnownCategory > category {
			oldAvg := db.GetAvgScore(year, category)
			newAvg := db.GetAvgScore(year, lastKnownCategory)
			newMax := db.GetMaxScore(year, lastKnownCategory)
			thisYearScore := baseScore

			convertedScore := thisYearScore * newAvg / oldAvg
			// Do not exceed the maximum obtainable score
			if convertedScore > newMax {
				convertedScore = newMax
			}
			baseScore = convertedScore
		}

		// Aging:
		// Most recent scores should weight more than past years ones
		agingFactor := 1 - 5/72*math.Log2(float64(lastKnownYear-year+1))
		baseScore = baseScore * agingFactor

		reelo = reelo + baseScore
		count++
	}

	reelo = reelo / float64(count)

	// Anti-Exploit:
	// To avoid the exploiting of a single partecipation, if the player has only
	// one result and it's in the most recent year: her/his REELO decays
	// TODO: do not use time.Now().Year(), recover a list of all known years
	if count == 1 && lastKnownYear == time.Now().Year() {
		reelo = reelo * ANTI_EXPLOIT
	}

	// No-partecipation penalty:
	// If the player hasn't partecipated in the most recent year, his REELO decaysa
	// TODO actually check if there's no result
	// if NO RESULT IN THE LAST YEAR {
	//    reelo = reelo * NP_PENALTY
	// }
	return reelo
}
