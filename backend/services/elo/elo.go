package elo

import (
	"context"
	"log"

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
// Variables names are chosen consistently with the formula
// provided by the scientific committee
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

// PseudoReelo calculates a basic version of a player's ELO.
// This score does not take aging, anti-exploit and category
// promotion into consideration.
func PseudoReelo(ctx context.Context, name, surname string, year int) error {
	var isParis bool
	db := rdb.NewDB()
	defer db.Close()

	log.Printf("----- INIZIO pseudo-Reelo per %v %v\n", name, surname)
	//### Steps from 1 to 5
	// There could be more than one category for a year,
	// this could happen in case of namesakes or international results
	log.Printf("[%v %v] Considero l'Anno: %v\n", name, surname, year)
	categories, err := db.Categories(ctx, name, surname, year)
	if err != nil {
		return err
	}

	for _, c := range categories {
		log.Printf("[%v %v] Categoria: %v\n", name, surname, c)
		isParis, err := db.IsResultFromParis(ctx, name, surname, year, c)
		if err != nil {
			return err
		}
		log.Printf("[%v %v] È parigi?: %v\n", name, surname, isParis)
	}

	for _, c := range categories {
		reelo, err := oneYearScore(ctx, name, surname, c, year, isParis)
		if err != nil {
			return err
		}
		// TODO: this should be in a service, not here
		err = db.UpdatePseudoReelo(ctx, name, surname, year, c, reelo)
		if err != nil {
			return err
		}
	}
	return nil
}

// Reelo calculates a player's ELO using a custom algorithm
func Reelo(ctx context.Context, name, surname string) (float64, error) {
	var reelo float64
	var sumOfWeights float64

	db := rdb.NewDB()
	defer db.Close()

	log.Printf("----- INIZIO Reelo per %v %v\n", name, surname)
	// Get some usefull values from db:
	//
	// A list of years in which the player has partecipated
	// It's used to iterate over the results as well as to check if the
	// anti-exploit mechanism should take effect.
	years, err := db.PlayerPartecipationYears(ctx, name, surname)
	if err != nil {
		return reelo, err
	}
	// The last category known in which the player has partecipated.
	// It's used to check for category promotion.
	lastKnownCategoryForPlayer, err := db.LastKnownCategoryForPlayer(name, surname)
	if err != nil {
		return reelo, err
	}
	// The last year known in which the player has partecipated.
	// It is used to calculate the aging factor and if the anti-exploit mechanism
	// should take effect.
	lastKnownYear, err := db.LastKnownYear()
	if err != nil {
		return reelo, err
	}

	// Every year a player has played should have a pseudo-Reelo, which needs
	// to be aged and promoted (if needed) to become a proper Reelo.
	// Finally we want to calculate a weighted average of
	// all those Reeloj and get the final score.
	for _, year := range years {
		// pseudo-Reelo is basically a glorified version of the K*E+D formula.
		pseudoReelo, err := db.PseudoReelo(ctx, name, surname, year)
		if err != nil {
			return reelo, err
		}

		// the category for the current year, used to check for category promotion
		category, err := db.Category(ctx, name, surname, year)
		if err != nil {
			return reelo, err
		}

		err = stepSix(&pseudoReelo, lastKnownCategoryForPlayer, category, year)
		if err != nil {
			return reelo, err
		}
		log.Printf("[%v %v OYS %v %v] Passo6: Promozione punteggio convertito %v\n", name, surname, year, category, pseudoReelo)

		stepSeven(&pseudoReelo, &sumOfWeights, lastKnownYear, year)

		reelo += pseudoReelo
	}

	log.Printf("[%v %v] Passo8: sommaReeloj/sommaPesi: %v / %v = %v\n", name, surname, reelo, sumOfWeights, reelo/sumOfWeights)
	stepEight(&reelo, sumOfWeights)

	stepNine(&reelo, years, lastKnownYear)
	stepTen(&reelo, years, lastKnownYear)

	log.Printf("[%v %v] Punteggio finale: %v\n", name, surname, reelo)
	return reelo, nil
}

// oneYearScore is used to calculate a baseScore using steps from 1 to 5.
// This baseScore (a.k.a. pseudo-Reelo) refers to a single year.
// This needs to be calculated for every year the player has played.
func oneYearScore(ctx context.Context, name, surname, category string,
	year int, isParis bool) (float64, error) {
	var baseScore float64
	db := rdb.NewDB()
	defer db.Close()

	// the first exercise a player is supposed to solve for the given category
	t, err := db.StartOfCategory(context.Background(), year, category)
	if err != nil {
		return baseScore, err
	}
	log.Printf("[%v %v OYS %v %v] primo esercizio della cat: %v\n", name, surname, year, category, t)

	// the last exercise a player is supposed to solve for the given category
	n, err := db.EndOfCategory(context.Background(), year, category)
	if err != nil {
		return baseScore, err
	}

	log.Printf("[%v %v OYS %v %v] ultimo esercizio della cat: %v\n", name, surname, year, category, n)

	// the maximum number of solvable exercises for the given category
	eMax := float64(n - t + 1)
	log.Printf("[%v %v OYS %v %v] eMax: %v\n", name, surname, year, category, eMax)

	maxScoreForCat, err := db.MaxScoreForCategory(context.Background(), year, category)
	if err != nil {
		return baseScore, err
	}
	// the maximum score obtainable in the given category
	dMax := float64(maxScoreForCat)
	log.Printf("[%v %v OYS %v %v] dMax: %v\n", name, surname, year, category, dMax)

	// the player's score for this year-category
	d, err := db.Score(name, surname, year, isParis)
	if err != nil {
		return baseScore, err
	}

	// the number of exercises solved by the player for this year-category
	exercises, err := db.Exercises(name, surname, year, isParis)
	if err != nil {
		return baseScore, err
	}
	e := float64(exercises)
	log.Printf("[%v %v OYS %v %v] numero di esercizi svolti dal giocatore: %v\n", name, surname, year, category, e)

	stepOne(&baseScore, e, d)
	log.Printf("[%v %v OYS %v %v] Passo1: punteggioBase = k*e+d = %v*%v+%v = %v\n", name, surname, year, category, exercisesCostant, e, d, baseScore)

	stepTwo(&baseScore, isParis)
	log.Printf("[%v %v OYS %v %v] Passo2: dopo parigi %v\n", name, surname, year, category, baseScore)

	stepThree(&baseScore, t, n, d, e, eMax, dMax)
	log.Printf("[%v %v OYS %v %v] Passo3: Omogeneizzazione finita, punteggio base: %v\n", name, surname, year, category, baseScore)

	err = stepFour(&baseScore, year)
	if err != nil {
		return baseScore, err
	}
	log.Printf("[%v %v OYS %v %v] Passo4: punteggio base dopo normalizzazione: %v\n", name, surname, year, category, baseScore)

	stepFive(&baseScore)
	log.Printf("[%v %v OYS %v %v] Passo5: dopo fattore moltiplicativo: %v\n", name, surname, year, category, baseScore)

	return baseScore, nil
}

func maxPseudoReelo(year int, category string) (float64, error) {
	var pseudoReelo float64
	db := rdb.NewDB()
	defer db.Close()

	t, err := db.StartOfCategory(context.Background(), year, category)
	if err != nil {
		return pseudoReelo, err
	}
	n, err := db.EndOfCategory(context.Background(), year, category)
	if err != nil {
		return pseudoReelo, err
	}

	// Here eMax correspond to e
	eMax := float64(n - t + 1)
	e := eMax
	// Here dMax correspond to d
	maxScoreForCat, err := db.MaxScoreForCategory(context.Background(), year, category)
	if err != nil {
		return pseudoReelo, err
	}
	// the maximum score obtainable in the given category
	dMax := float64(maxScoreForCat)
	d := dMax

	log.Printf("Passo6: Promozione, calcolo il massimo Reelo ottenibile nel %v categoria %v\n", year, category)
	stepOne(&pseudoReelo, e, d)
	stepTwo(&pseudoReelo, true)
	stepThree(&pseudoReelo, t, n, d, e, eMax, dMax)
	stepFour(&pseudoReelo, year)
	stepFive(&pseudoReelo)

	return pseudoReelo, nil
}

func contains(array []int, item int) bool {
	for _, e := range array {
		if e == item {
			return true
		}
	}
	return false
}
