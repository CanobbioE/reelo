package elo

import (
	"log"
	"math"

	rdb "github.com/CanobbioE/reelo/backend/db"
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
func stepThree(baseScore *float64, t, n int, d, e, eMax, dMax float64) {
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

//### 4. Score normailzation:
// Scores are normalized to the average of averages of this year's categories
func stepFour(baseScore *float64, year int) error {
	db := rdb.NewDB()
	defer db.Close()
	avgCatScore, err := db.AvgScoresOfCategories(year, exercisesCostant)
	if err != nil {
		return err
	}
	log.Printf("Passo4: media delle medie delle categorie: %v\n", avgCatScore)
	log.Printf("Passo4: punteggioBase / mediaDelleMedie = %v / %v = %v", *baseScore, avgCatScore, *baseScore/avgCatScore)
	*baseScore = *baseScore / avgCatScore
	return nil
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
	category string, year int) error {
	db := rdb.NewDB()
	defer db.Close()

	log.Printf("Passo6: Promozione ultima categoria conosciuta: %v\n", lastKnownCategoryForPlayer)
	log.Printf("Passo6: Promozione categoria attuale: %v\n", category)
	if categoryFromString(lastKnownCategoryForPlayer) > categoryFromString(category) {
		oldAvg, err := db.AvgPseudoReelo(year, category)
		if err != nil {
			return err
		}
		log.Printf("Passo6: Promozione vecchia media: %v\n", oldAvg)

		newAvg, err := db.AvgPseudoReelo(year, lastKnownCategoryForPlayer)
		if err != nil {
			return err
		}
		// this prevents crash in case the ranking file for the
		// year-lastKnownCategory has not been uploaded yet
		if newAvg < 0 {
			return nil
		}
		log.Printf("Passo6: Promozione nuova media: %v\n", newAvg)

		newMax, err := maxPseudoReelo(year, lastKnownCategoryForPlayer)
		if err != nil {
			return err
		}
		// this prevents crash in case the ranking file for the
		// year-lastKnownCategory has not been uploaded yet
		if newMax < 0 {
			return nil
		}
		log.Printf("Passo6: Promozione nuovo massimo: %v\n", newMax)

		thisYearScore := *baseScore

		convertedScore := thisYearScore * newAvg / oldAvg
		// Do not exceed the maximum obtainable score
		if convertedScore > newMax {
			log.Printf("Passo6: Promozione superato massimo!: %v > %v\n", convertedScore, newMax)
			convertedScore = newMax
		}
		//  Do not drop below the originalScore
		if convertedScore < thisYearScore {
			log.Printf("Passo6: Promozione superato il minimo!: %v < %v\n", convertedScore, thisYearScore)
			convertedScore = thisYearScore
		}
		*baseScore = convertedScore
	}
	return nil
}

//### 7. Aging:
// Most recent scores should weight more than the ones from past years
func stepSeven(baseScore, sumOfWeights *float64, lastKnownYear, year int) {
	agingFactor := 1 - float64(5)/72*math.Log2(float64(lastKnownYear-year+1))
	log.Printf("Passo7: Fattore invecchiamento 1 - 5/72*log2(A-Ai+1) = 1 - 5/72 * log2(%v-%v+1): %v\n", lastKnownYear, year, agingFactor)
	*baseScore *= agingFactor
	*sumOfWeights += agingFactor
	log.Printf("Passo7: punteggioBasse dopo invecchiamento: %v\n", *baseScore)
}

//### 8. Average:
// Since reelo is already the sum of weighted scores we just divide
// by the sum of the weights
func stepEight(baseScore *float64, sumOfWeights float64) {
	*baseScore = *baseScore / sumOfWeights
}

//### 9. Anti-Exploit:
// To avoid single year partecipation's exploit: if the player has only
// one result and it's in the most recent year, then her/his REELO is worth less
func stepNine(baseScore *float64, years []int, lastKnownYear int) {
	if len(years) == 1 && lastKnownYear == years[0] {
		log.Printf("Passo9: antiExploit: %v, reelo = reelo * antiExploit = %v \n", antiExploit, *baseScore*antiExploit)
		*baseScore *= antiExploit
	}
}

//### 10. No-partecipation penalty:
// If the player didn't partecipate in the most recent year, his REELO is worth less
func stepTen(baseScore *float64, years []int, lastKnownYear int) {
	if !contains(years, lastKnownYear) {
		log.Printf("Passo10: penalità non partecipazione: %v * %v = %v\n", *baseScore, noPartecipationPenalty, *baseScore*noPartecipationPenalty)
		*baseScore *= noPartecipationPenalty
	}
}
