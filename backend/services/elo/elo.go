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
	log.Printf("[%v %v]ultimaCat: %v ultimoAnno: %v anniPartecipa: %v\n", name, surname, lastKnownCategoryForPlayer, lastKnownYear, partecipationYears)

	//### Steps from 1 to 7
	for _, year := range partecipationYears {
		// There could be more than one category for a year,
		// this could happen in case of namesakes or international results

		log.Printf("[%v %v] Considero l'Anno: %v\n", name, surname, year)
		categories, err := db.Categories(ctx, name, surname, year)
		if err != nil {
			return reelo, err
		}

		for _, c := range categories {
			log.Printf("[%v %v] Categoria: %v\n", name, surname, c)
			isParis, err := db.IsResultFromParis(ctx, name, surname, year, c)
			if err != nil {
				return reelo, err
			}
			log.Printf("[%v %v] È parigi?: %v\n", name, surname, isParis)

			weight, err := oneYearScore(ctx,
				name, surname, lastKnownCategoryForPlayer, c,
				year, lastKnownYear, &reelo, isParis)
			if err != nil {
				return reelo, err
			}
			log.Printf("[%v %v] Aggiungo ai pesi %v il nuovo peso: %v\n", name, surname, weights, weight)
			weights = append(weights, weight)
		}
	}

	//### 8. Average:
	// Since reelo is already the sum of weighted scores we just divide
	// by the sum of the weights
	var sumOfWeights float64
	for _, w := range weights {
		sumOfWeights += w
	}

	log.Printf("[%v %v] Passo8: reelo/sommaPesi: %v / %v = %v\n", name, surname, reelo, sumOfWeights, reelo/sumOfWeights)
	reelo = reelo / sumOfWeights

	//### 9. Anti-Exploit:
	// To avoid single year partecipation's exploit: if the player has only
	// one result and it's in the most recent year, then her/his REELO is worth less
	if len(partecipationYears) == 1 && lastKnownYear == partecipationYears[0] {
		log.Printf("[%v %v] Passo9: antiExploit: %v, reelo = reelo * antiExploit = %v \n", name, surname, antiExploit, reelo*antiExploit)
		reelo = reelo * antiExploit
	}

	//### 10. No-partecipation penalty:
	// If the player didn't partecipate in the most recent year, his REELO is worth less
	if !contains(partecipationYears, lastKnownYear) {
		log.Printf("[%v %v] Passo10: penalità non partecipazione: %v * %v = %v\n", name, surname, reelo, noPartecipationPenalty, reelo*noPartecipationPenalty)
		reelo = reelo * noPartecipationPenalty
	}

	log.Printf("[%v %v] Punteggio finale finale: %v\n", name, surname, reelo)
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
	log.Printf("[%v %v OYS %v %v] primo esercizio della cat: %v\n", name, surname, year, category, t)

	n, err := db.EndOfCategory(context.Background(), year, category)
	if err != nil {
		return 0, err
	}

	log.Printf("[%v %v OYS %v %v] ultimo esercizio della cat: %v\n", name, surname, year, category, n)

	eMax := float64(n - t + 1)
	log.Printf("[%v %v OYS %v %v] eMax: %v\n", name, surname, year, category, eMax)
	maxScoreForCat, err := db.MaxScoreForCategory(context.Background(), year, category)
	if err != nil {
		return 0, err
	}

	dMax := float64(maxScoreForCat)
	log.Printf("[%v %v OYS %v %v] dMax: %v\n", name, surname, year, category, dMax)
	d, err := db.Score(name, surname, year, isParis)
	if err != nil {
		return 0, err
	}
	log.Printf("[%v %v OYS %v %v] punteggio ottenuto dal giocatore: %v\n", name, surname, year, category, eMax)

	exercises, err := db.Exercises(name, surname, year, isParis)
	if err != nil {
		return 0, err
	}
	e := float64(exercises)
	log.Printf("[%v %v OYS %v %v] numero di esercizi svolti dal giocatore: %v\n", name, surname, year, category, e)

	//### 1. Base score:
	// Is the sum of the difficulty D and the number of completed exercises
	baseScore := exercisesCostant*e + d

	log.Printf("[%v %v OYS %v %v] Passo1: punteggioBase = k*e+d = %v*%v+%v = %v\n", name, surname, year, category, exercisesCostant, e, d, baseScore)

	//### 2. International results:
	// If the result is from paris let's multiply for pFinal
	if isParis {
		baseScore = baseScore * pFinal
		log.Printf("[%v %v OYS %v %v] Passo2: dopo parigi %v\n", name, surname, year, category, baseScore)
	}

	//### 3. Category homogenization:
	categoriesHomogenization(&baseScore, t, n, d, e, eMax, dMax)
	log.Printf("[%v %v OYS %v %v] Passo3: Omogeneizzazione finita, punteggio base: %v\n", name, surname, year, category, baseScore)

	//### 4. Score normailzation:
	// Scores are normalized to the average of averages of this year's categories
	avgCatScore, err := db.AvgScoresOfCategories(year)
	if err != nil {
		return 0, err
	}
	log.Printf("[%v %v OYS %v %v] Passo4: media delle medie delle categorie: %v\n", name, surname, year, category, avgCatScore)
	baseScore = baseScore / avgCatScore
	log.Printf("[%v %v OYS %v %v] Passo4: punteggio base dopo normalizzazione: %v\n", name, surname, year, category, baseScore)

	//### 5. Multiplicative factor:
	// just to have a big nice number let's multiply for a costant
	baseScore = baseScore * multiplicativeFactor
	log.Printf("[%v %v OYS %v %v] Passo5: dopo fattore moltiplicativo: %v\n", name, surname, year, category, baseScore)

	//### 6. Category promotion:
	categoryPromotion(&baseScore, lastKnownCategoryForPlayer, category, year)
	log.Printf("[%v %v OYS %v %v] Passo6: Promozione punteggio convertito %v\n", name, surname, year, category, baseScore)

	//### 7. Aging:
	// Most recent scores should weight more than past years ones
	agingFactor := 1 - float64(5)/72*math.Log2(float64(lastKnownYear-year+1))
	log.Printf("[%v %v OYS %v %v] Passo7: Fattore invecchiamento 1 - 5/72*log2(A-Ai+1) = 1 - 5/72 * log2(%v-%v+1): %v\n", name, surname, year, category, lastKnownYear, year, agingFactor)
	baseScore = baseScore * agingFactor
	log.Printf("[%v %v OYS %v %v] Passo7: punteggioBasse dopo invecchiamento: %v\n", name, surname, year, category, baseScore)

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
		log.Printf("Passo3: Omogeneizzazione fattore errore: %v\n", errorFactor)
		difficultyFactor := float64(i) / float64(n+1)
		log.Printf("Passo3: Omogeneizzazione fattore difficoltà: %v\n", difficultyFactor)
		nonResolutionProbability := 1 - difficultyFactor*errorFactor
		log.Printf("Passo3: Omogeneizzazione prob di non risoluzione: %v\n", nonResolutionProbability)
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

	log.Printf("Passo6: Promozione ultima categoria conosciuta: %v\n", lastKnownCategoryForPlayer)
	log.Printf("Passo6: Promozione categoria attuale: %v\n", category)
	if categoryFromString(lastKnownCategoryForPlayer) > categoryFromString(category) {
		oldAvg, err := db.AvgScore(year, category)
		if err != nil {
			return err
		}
		log.Printf("Passo6: Promozione vecchia media: %v\n", oldAvg)
		newAvg, err := db.AvgScore(year, lastKnownCategoryForPlayer)
		if err != nil {
			return err
		}
		log.Printf("Passo6: Promozione nuova media: %v\n", newAvg)
		newMax, err := db.MaxScore(year, lastKnownCategoryForPlayer)
		if err != nil {
			return err
		}
		log.Printf("Passo6: Promozione nuovo massimo: %v\n", newMax)

		thisYearScore := *baseScore

		convertedScore := thisYearScore * newAvg / oldAvg
		// Do not exceed the maximum obtainable score
		if convertedScore > newMax {
			log.Printf("Passo6: Promozione superato massimo!: %v > %v\n", convertedScore, newMax)
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
