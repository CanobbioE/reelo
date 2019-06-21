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

const debug = true

// InitCostants retrieves all the costants in the database, if anything goes wrong
// it will fallback to the hardcoded values
func InitCostants() {
	db := rdb.NewDB()
	defer db.Close()

	c, err := db.ReeloCostants()
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
func Reelo(ctx context.Context, name, surname string) (float64, error) {
	var reelo float64
	db := rdb.NewDB()
	defer db.Close()

	lastKnownCategoryForPlayer, err := db.LastKnownCategoryForPlayer(name, surname)
	if err != nil {
		return reelo, err
	}
	lastKnownYear, err := db.LastKnownYear()
	if err != nil {
		return reelo, err
	}
	partecipationYears, err := db.PlayerPartecipationYears(ctx, name, surname)
	if err != nil {
		return reelo, err
	}

	//### Steps from 1 to 7
	for _, year := range partecipationYears {
		// There could be more than one category for a year,
		// this could happen for namesakes or when we have international results
		categories, err := db.Categories(ctx, name, surname, year)
		if err != nil {
			return reelo, err
		}
		for _, c := range categories {
			isParis, err := db.IsResultFromParis(ctx, name, surname, year, c)
			if err != nil {
				return reelo, err
			}
			err = oneYearScore(ctx,
				name, surname, lastKnownCategoryForPlayer, c,
				year, lastKnownYear, &reelo, isParis)
			if err != nil {
				return reelo, err
			}
		}
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

	return reelo, nil
}

func oneYearScore(ctx context.Context,
	name, surname, lastKnownCategoryForPlayer, category string,
	year, lastKnownYear int, reelo *float64, isParis bool) error {
	db := rdb.NewDB()
	defer db.Close()

	// Variables names are chosen accordingly to the formula
	// provided by the scientific committee
	t := StartOfCategory(year, category)
	n := EndOfCategory(year, category)
	eMax := float64(t - n + 1)
	dMax := float64(MaxScoreForCategory(year, category))
	d, err := db.Score(name, surname, year, isParis)
	if err != nil {
		return err
	}
	exercises, err := db.Exercises(name, surname, year, isParis)
	if err != nil {
		return err
	}
	e := float64(exercises)
	if err != nil {
		return err
	}

	//### 1. Base score:
	// Is the sum of the difficulty D and the number of completed exercises
	baseScore := d + e*exercisesCostant

	//### 2. International results:
	// If the result is from paris let's multiply for pFinal
	if isParis {
		baseScore = baseScore * pFinal
	}

	//### 3. Category homogenization:
	categoriesHomogenization(&baseScore, t, n, d, e, eMax, dMax)

	//### 4. Score normailzation:
	// Scores are normalized to the average of averages of this year's categories
	avgCatScore, err := db.AvgScoresOfCategories(year)
	if err != nil {
		return err
	}
	baseScore = baseScore / avgCatScore

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
	return nil
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
	year int) error {
	db := rdb.NewDB()
	defer db.Close()
	if categoryFromString(lastKnownCategoryForPlayer) > categoryFromString(category) {
		oldAvg, err := db.AvgScore(year, category)
		if err != nil {
			return err
		}
		newAvg, err := db.AvgScore(year, lastKnownCategoryForPlayer)
		if err != nil {
			return err
		}
		newMax, err := db.MaxScore(year, lastKnownCategoryForPlayer)
		if err != nil {
			return err
		}
		thisYearScore := *baseScore

		convertedScore := thisYearScore * newAvg / oldAvg
		// Do not exceed the maximum obtainable score
		if convertedScore > newMax {
			convertedScore = newMax
		}
		*baseScore = convertedScore
	}
	return nil
}

func contains(array []int, item int) bool {
	for _, e := range array {
		if e == item {
			return true
		}
	}
	return false
}
