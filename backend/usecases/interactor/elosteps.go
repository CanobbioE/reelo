package interactor

import (
	"math"

	"github.com/CanobbioE/reelo/backend/pkg/category"
)

//### 1. Base score:
// Is the sum of the difficulty D and the number of completed exercises
func stepOne(baseScore *float64, e, d float64) {
	*baseScore = exercisesCostant*e + d
}

//### 2. International results:
// If the result is from paris let's multiply for pFinal
func stepTwo(baseScore *float64, isParis bool) {
	if isParis {
		*baseScore *= pFinal
	}
}

//### 3. Categories homogenization:
// For each exercises a player is not supposed to solve we calculate
// her/his probability of solving it.
// To the base score we should add:
// (K + i) * {1 - [i/N+1] * [1 - (K*e+d)/(KeMax+dMax)]}
//
// Where:
//
// exercisesCostant = K
// errorFactor = 1 - [(K*e+d) / (K*eMax+dMax)]
// difficultyFactor = i/(n+1)
// nonResolutionProbability = 1 - difficultyFactor*errorFactor
func stepThree(baseScore *float64, t, n int, d, e, eMax, dMax float64) {
	for i := 1; i <= t-1; i++ {
		errorFactor := float64(1 - (exercisesCostant*e+d)/(exercisesCostant*eMax+dMax))
		difficultyFactor := float64(i) / float64(n+1)
		nonResolutionProbability := 1 - difficultyFactor*errorFactor
		*baseScore += (exercisesCostant + float64(i)) * nonResolutionProbability
	}
}

//### 4. Score normalization:
// Scores are normalized to the average of averages of this year's categories
func stepFour(baseScore *float64, year int, avgCatScore float64) {
	*baseScore = *baseScore / avgCatScore
}

//### 5. Multiplicative factor:
// just to have a big nice number let's multiply for a costant
func stepFive(baseScore *float64) {
	*baseScore *= multiplicativeFactor
}

//### 6. Category promotion:
// If in the year we are using to calculate the ELO,
// the player's category is inferior to the category she/he
// has played most recently, then we convert this year's score
// to the most recent category
func stepSix(baseScore *float64, lastKnownCategoryForPlayer,
	cat string, year int,
	newAvg, oldAvg, newMax float64) {

	if category.FromString(lastKnownCategoryForPlayer) > category.FromString(cat) {

		// this prevents crash in case the ranking file for the
		// year-lastKnownCategory has not been uploaded yet
		if newAvg < 0 {
			return
		}

		// this prevents crash in case the ranking file for the
		// year-lastKnownCategory has not been uploaded yet
		if newMax < 0 {
			return
		}

		thisYearScore := *baseScore

		convertedScore := thisYearScore * newAvg / oldAvg
		// Do not exceed the maximum obtainable score
		if convertedScore > newMax {
			convertedScore = newMax
		}
		//  Do not drop below the originalScore
		if convertedScore < thisYearScore {
			convertedScore = thisYearScore
		}
		*baseScore = convertedScore
	}
}

//### 7. Aging:
// Most recent scores should weight more than the ones from past years
func stepSeven(baseScore, sumOfWeights *float64, lastKnownYear, year int) {
	agingFactor := 1 - float64(5)/72*math.Log2(float64(lastKnownYear-year+1))
	*baseScore *= agingFactor
	*sumOfWeights += agingFactor
}

//### 8. Average:
// Since reelo is already the sum of weighted scores we just divide
// by the sum of the weights
func stepEight(baseScore *float64, sumOfWeights float64) {
	*baseScore = *baseScore / sumOfWeights
}

//### 9. Anti-Exploit:
// To avoid single year participation's exploit: if the player has only
// one result and it's in the most recent year, then her/his REELO is worth less
func stepNine(baseScore *float64, years []int, lastKnownYear int) {
	if len(years) == 1 && lastKnownYear == years[0] {
		*baseScore *= antiExploit
	}
}

//### 10. No-participation penalty:
// If the player didn't participate in the most recent year, his REELO is worth less
func stepTen(baseScore *float64, years []int, lastKnownYear int) {
	if !contains(years, lastKnownYear) {
		*baseScore *= noParticipationPenalty
	}
}
