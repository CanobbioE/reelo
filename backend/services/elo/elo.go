/*
Package elo defines the custom algorithm, and all the useful function used
to calculate the REELO for a given user.
Disclaimer: comments that starts with three '#' represents the steps
to calculate the REELO taken directly from the given documentation.
*/
package elo

import (
	"context"
	"log"
	"math"

	rdb "github.com/CanobbioE/reelo/backend/db"
)

var (
	startingYear           = 2002
	exercisesCostant       = 20.0
	pFinal                 = 1.5
	multiplicativeFactor   = 10000.0
	antiExploit            = 0.9
	noPartecipationPenalty = 0.9
)

// InitCostants retrieves all the costants in the database, if anything goes wrong
// it will fallback to the hardcoded values
func InitCostants() {
	db := rdb.NewDB()
	defer db.Close()

	c, err := db.GetReeloCostants()
	if err != nil {
		log.Printf("Error initializing costants: %v", err)
		log.Println("Falling back to the hardcoded configuration")
		return
	}
	startingYear = c.StartingYear
	exercisesCostant = c.ExercisesCostant
	pFinal = c.PFinal
	multiplicativeFactor = c.MultiplicativeFactor
	antiExploit = c.AntiExploit
	noPartecipationPenalty = c.NoPartecipationPenalty
}

// Reelo returns the points for a given user calculated with a custom algorithm.
func Reelo(ctx context.Context, name, surname string) (reelo float64) {
	db := rdb.NewDB()
	defer db.Close()

	lastKnownCategoryForPlayer := db.GetLastKnownCategoryForPlayer(name, surname)
	lastKnownYear := db.GetLastKnownYear()
	partecipationYears := db.GetPlayerPartecipationYears(ctx, name, surname)

	//### Steps from 1 to 7
	for _, year := range partecipationYears {
		oneYearScore(ctx,
			name, surname, lastKnownCategoryForPlayer,
			year, lastKnownYear, &reelo)
	}

	//### 8. Average:
	// Since reelo is already the sum of weighted scores we just divide
	// by the number of years we know the player has played
	reelo = reelo / float64(len(partecipationYears))

	//### 9. Anti-Exploit:
	// To avoid the exploitation of a single partecipation: if the player has only
	// one result and it's in the most recent year, then her/his REELO decays
	if len(partecipationYears) == 1 && lastKnownYear == partecipationYears[0] {
		reelo = reelo * antiExploit
	}

	//### 10. No-partecipation penalty:
	// If the player didn't partecipate in the most recent year, his REELO decays
	if !contains(partecipationYears, lastKnownYear) {
		reelo = reelo * noPartecipationPenalty
	}
	return reelo
}

func oneYearScore(ctx context.Context,
	name, surname, lastKnownCategoryForPlayer string,
	year, lastKnownYear int, reelo *float64) {
	db := rdb.NewDB()
	defer db.Close()

	// Variables names are chosen accordingly to the formula
	// provided by the scientific committee
	category := db.GetCategory(name, surname, year)

	t := StartOfCategory(year, category)
	n := EndOfCategory(year, category)
	eMax := float64(t - n + 1)
	dMax := float64(MaxScoreForCategory(year, category))
	d := db.GetScore(name, surname, year)
	e := float64(db.GetExercises(name, surname, year))

	//### 1. Base score:
	// Is the sum of the difficulty D and the number of completed exercises
	baseScore := d + e*exercisesCostant

	//### 2. International results:
	// If the result is from paris let's multiply for pFinal
	if db.IsResultFromParis(name, surname, year, category) {
		baseScore = baseScore * pFinal
	}

	//### 3. Category homogenization:
	categoriesHomogenization(&baseScore, t, n, d, e, eMax, dMax)

	//### 4. Score normailzation:
	// Scores ar normalized to the average of averages of this year's categories
	baseScore = baseScore / db.GetAvgScoresOfCategories(year)

	//### 5. Multiplicative factor:
	// just to have a big nice number let's multiply for a costant
	baseScore = baseScore * multiplicativeFactor

	//### 6. Category promotion:
	categoryPromotion(&baseScore, lastKnownCategoryForPlayer, category, year)

	//### 7. Aging:
	// Most recent scores should weight more than past years ones
	agingFactor := 1 - 5/72*math.Log2(float64(lastKnownYear-year+1))
	baseScore = baseScore * agingFactor

	*reelo = *reelo + baseScore
}

//### 3. Categories homogenization:
// For each exercises a player is not supposed to solve we calculate
// her/his probabilty of solving it.
func categoriesHomogenization(baseScore *float64, t, n int, d, e, eMax, dMax float64) {
	for i := 1; i <= t-1; i++ {
		errorFactor := float64(1-d+e*exercisesCostant) / (exercisesCostant*eMax + dMax)
		difficultyFactor := float64(i) / float64(n+1)
		nonResolutionProbability := 1 - errorFactor*difficultyFactor
		*baseScore += (exercisesCostant + float64(i)) * nonResolutionProbability
	}
}

//### 6. Category promotion:
// If this year the player's category is inferior to the category she/he
// has played most recently, then we convert this year score
// to the most recent category
func categoryPromotion(baseScore *float64,
	lastKnownCategoryForPlayer, category string,
	year int) {
	db := rdb.NewDB()
	defer db.Close()
	if categoryFromString(lastKnownCategoryForPlayer) > categoryFromString(category) {
		oldAvg := db.GetAvgScore(year, category)
		newAvg := db.GetAvgScore(year, lastKnownCategoryForPlayer)
		newMax := db.GetMaxScore(year, lastKnownCategoryForPlayer)
		thisYearScore := *baseScore

		convertedScore := thisYearScore * newAvg / oldAvg
		// Do not exceed the maximum obtainable score
		if convertedScore > newMax {
			convertedScore = newMax
		}
		*baseScore = convertedScore
	}
}

func contains(array []int, item int) bool {
	for _, e := range array {
		if e == item {
			return true
		}
	}
	return false
}
