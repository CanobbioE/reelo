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

// InitCostants retrieves the costants in the database, if anything goes wrong
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

// Reelo calculates a player's ELO using a custom algorithm
func Reelo(ctx context.Context, name, surname string) (float64, error) {
	var reelo float64
	var weights []float64
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
	log.Printf("----- INIZIO per %v %v\n", name, surname)
	log.Printf("[%v %v]ultimaCat: %v\tultimoAnno :%v\tanniPartecipato :%v\n", name, surname, lastKnownCategoryForPlayer, lastKnownYear, partecipationYears)

	//### Steps from 1 to 7
	for _, year := range partecipationYears {
		// There could be more than one category for a year,
		// this could happen in case of namesakes or international results

		log.Printf("[%v %v] Anno: %v\n", name, surname, year)
		categories, err := db.Categories(ctx, name, surname, year)
		if err != nil {
			return reelo, err
		}
		log.Printf("[%v %v] Categorie: %v\n", name, surname, categories)

		for _, c := range categories {
			isParis, err := db.IsResultFromParis(ctx, name, surname, year, c)
			if err != nil {
				return reelo, err
			}
			log.Printf("[%v %v] Parigi?: %v\n", name, surname, isParis)

			weight, err := oneYearScore(ctx,
				name, surname, lastKnownCategoryForPlayer, c,
				year, lastKnownYear, &reelo, isParis)
			if err != nil {
				return reelo, err
			}
			weights = append(weights, weight)
			log.Printf("[%v %v] Pesi: %v\n", name, surname, weights)
		}
	}

	//### 8. Average:
	// Since reelo is already the sum of weighted scores we just divide
	// by the sum of the weights
	var sumOfWeights float64
	for _, w := range weights {
		sumOfWeights += w
	}

	log.Printf("[%v %v] sommaPesi: %v\treelo: %v\treelo/sumPesi: %v\n", name, surname, sumOfWeights, reelo, reelo/sumOfWeights)
	reelo = reelo / sumOfWeights

	//### 9. Anti-Exploit:
	// To avoid single year partecipation's exploit: if the player has only
	// one result and it's in the most recent year, then her/his REELO is worth less
	if len(partecipationYears) == 1 && lastKnownYear == partecipationYears[0] {
		reelo = reelo * antiExploit
		log.Printf("[%v %v] Dopo antiExploit: %v\n", name, surname, reelo)
	}

	//### 10. No-partecipation penalty:
	// If the player didn't partecipate in the most recent year, his REELO is worth less
	if !contains(partecipationYears, lastKnownYear) {
		reelo = reelo * noPartecipationPenalty
		log.Printf("[%v %v] Dopo noPartecipa:%v\n", name, surname, reelo)
	}

	log.Printf("[%v %v] Finale: %v\n", name, surname, reelo)
	return reelo, nil
}

func oneYearScore(ctx context.Context,
	name, surname, lastKnownCategoryForPlayer, category string,
	year, lastKnownYear int, reelo *float64, isParis bool) (float64, error) {
	db := rdb.NewDB()
	defer db.Close()

	// Variables names are chosen consistently with the formula
	// provided by the scientific committee
	t, err := db.StartOfCategory(context.Background(), year, category)
	if err != nil {
		return 0, err
	}
	log.Printf("[%v %v OYS %v %v] StartOfCategory: %v\n", name, surname, year, category, t)

	n, err := db.EndOfCategory(context.Background(), year, category)
	if err != nil {
		return 0, err
	}

	log.Printf("[%v %v OYS %v %v] EndOfCategory: %v\n", name, surname, year, category, n)

	eMax := float64(n - t + 1)
	log.Printf("[%v %v OYS %v %v] eMax: %v\n", name, surname, year, category, eMax)
	maxScoreForCat, err := db.MaxScoreForCategory(context.Background(), year, category)
	if err != nil {
		return 0, err
	}
	log.Printf("[%v %v OYS %v %v] MaxScoreForCategory: %v\n", name, surname, year, category, maxScoreForCat)

	dMax := float64(maxScoreForCat)
	d, err := db.Score(name, surname, year, isParis)
	if err != nil {
		return 0, err
	}
	log.Printf("[%v %v OYS %v %v] score: %v\n", name, surname, year, category, eMax)

	exercises, err := db.Exercises(name, surname, year, isParis)
	if err != nil {
		return 0, err
	}
	e := float64(exercises)
	log.Printf("[%v %v OYS %v %v] e-exercises: %v-%v\n", name, surname, year, category, e, exercises)

	//### 1. Base score:
	// Is the sum of the difficulty D and the number of completed exercises
	baseScore := exercisesCostant*e + d

	log.Printf("[%v %v OYS %v %v] baseScore: %v\n", name, surname, year, category, baseScore)

	//### 2. International results:
	// If the result is from paris let's multiply for pFinal
	if isParis {
		baseScore = baseScore * pFinal
		log.Printf("[%v %v OYS %v %v] baseScore dopo parigi: %v\n", name, surname, year, category, baseScore)
	}

	//### 3. Category homogenization:
	categoriesHomogenization(&baseScore, t, n, d, e, eMax, dMax)
	log.Printf("[%v %v OYS %v %v] baseScore dopo catHomo: %v\n", name, surname, year, category, baseScore)

	//### 4. Score normailzation:
	// Scores are normalized to the average of averages of this year's categories
	avgCatScore, err := db.AvgScoresOfCategories(year)
	if err != nil {
		return 0, err
	}

	log.Printf("[%v %v OYS %v %v] avgCatScore: %v\n", name, surname, year, category, avgCatScore)
	baseScore = baseScore / avgCatScore
	log.Printf("[%v %v OYS %v %v] baseScore dopo avg: %v\n", name, surname, year, category, baseScore)

	//### 5. Multiplicative factor:
	// just to have a big nice number let's multiply for a costant
	baseScore = baseScore * multiplicativeFactor
	log.Printf("[%v %v OYS %v %v] baseScore dopo mult: %v\n", name, surname, year, category, baseScore)

	//### 6. Category promotion:
	categoryPromotion(&baseScore, lastKnownCategoryForPlayer, category, year)
	log.Printf("[%v %v OYS %v %v] baseScore dopo promotion: %v\n", name, surname, year, category, baseScore)

	//### 7. Aging:
	// Most recent scores should weight more than past years ones
	agingFactor := 1 - float64(5)/72*math.Log2(float64(lastKnownYear-year+1))
	log.Printf("[%v %v OYS %v %v] aging: %v\n", name, surname, year, category, agingFactor)
	baseScore = baseScore * agingFactor
	log.Printf("[%v %v OYS %v %v] baseScore dopo aging: %v\n", name, surname, year, category, baseScore)

	*reelo = *reelo + baseScore
	return agingFactor, nil
}

//### 3. Categories homogenization:
// For each exercises a player is not supposed to solve we calculate
// her/his probabilty of solving it.
// To the base score we should add:
// (K + i) * {1 - [i/N+1] * [1 - (K*e+d)/(KeMax+dMax)]}
//
// Where:
//
// exercisesCostant = K
// errorFactor = 1 - [(K*e+d) / (K*eMax+dMax)]
// difficultyFactor = i/(n+1)
// nonResolutionProbability = 1 - difficultyFactor*errorFactor
func categoriesHomogenization(baseScore *float64, t, n int, d, e, eMax, dMax float64) {
	for i := 1; i <= t-1; i++ {
		errorFactor := float64(1 - (exercisesCostant*e+d)/(exercisesCostant*eMax+dMax))
		difficultyFactor := float64(i) / float64(n+1)
		nonResolutionProbability := 1 - difficultyFactor*errorFactor
		*baseScore += (exercisesCostant + float64(i)) * nonResolutionProbability
	}
}

//### 6. Category promotion:
// If in the year we are using to calculate the ELO,
// the player's category is inferior to the category she/he
// has played most recently, then we convert this year's score
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
