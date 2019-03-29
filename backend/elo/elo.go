package elo

import (
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
func Reelo(name, surname string) (reelo int) {
	// TODO: Reelo should not make DB calls
	/* TODO: transalte this to code
	// lastKnownCategory
	// lastKnownYear
	count = 0


	for _, year := range allYears() {
		cat = getCat(name, surname, year)
		T = startOfCategory(cat, year)
		N = endOfCategory(cat, year)
		eMax = T-N +1
		dMax = maxScoreForCategory(cat, year)
		D = sum(Di)
		E = numExSolved()

		baseScore =  D + E * K_EXERCISES

		if FINALE INTERNAZIONALE {
			baseScore = baseScore * P_FINAL
		}

		for i = 1; i <= T-1; i++ {
			// TODO ask if error factor can be calculated using baseScore (prob parigi)
			errorFactor = 1 - D + E * K_EXERCISES / (K_EXERCISES * eMax + dMax)
			difficultyFactor = i/(N+1)
			nonResolutionProbability = 1 - errorFactor * difficultyFactor)
			baseScore += (K_EXERCISES + i) * nonResolutionProbability
		}

		baseScore = baseScore / avgScoresOfCategories(year)
		baseScore = baseScore * MULTIPLICATIVE_FACTOR

		if lastKnownCategory > cat { // passaggio di categoria
			oldAvg = avgOfScores(year, cat)
			newAvg = avgOfScores(year, lastKnownCat)
			newMax = maxOfScores(year, lastKnownCat)
			scoreOfYear = baseScore

			convertedScore = scoreOfYear * newAvg / oldAvg
			if convertedScore > newMax {
				convertedScore = newMax
			}

			baseScore = convertedScore
		}

		agingFactor = 1 - 5/72 * math.Log2(lastKnownYear-year+1)
		baseScore = baseScore * agingFactor

		reelo = reelo + baseScore
		count++
	}

	reelo = reelo / count

	if count == 1 && lastKnwonYear == time.Now().Year() {
		reelo = reelo * ANTI_EXPLOIT
	}

	if NO RESULT INT THE LAST YEAR {
		reelo = reelo * NP_PENALTY
	}
	*/
	return reelo
}
