package interactor

import (
	"context"
	"log"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
	"github.com/CanobbioE/reelo/backend/utils/category"
)

var (
	startingYear           = 2002
	exercisesCostant       = 20.0
	pFinal                 = 1.5
	multiplicativeFactor   = 10000.0
	antiExploit            = 0.9
	noParticipationPenalty = 0.9
)

// InitCostants retrieves the costants in the database, if anything goes wrong
// it will fallback to the hardcoded values
// Variables names are chosen consistently with the formula
// provided by the scientific committee
func (i *Interactor) InitCostants() {

	c, err := i.CostantsRepository.FindAll(context.Background())
	if err != nil {
		i.Logger.Log("Error initializing costants: %v", err)
		i.Logger.Log("Falling back to the hardcoded configuration\n")
		return
	}
	startingYear = c.StartingYear
	exercisesCostant = c.ExercisesCostant
	pFinal = c.PFinal
	multiplicativeFactor = c.MultiplicativeFactor
	antiExploit = c.AntiExploit
	noParticipationPenalty = c.NoParticipationPenalty
}

// PseudoReelo calculates a basic version of a player's ELO.
// This score does not take aging, anti-exploit and category
// promotion into consideration.
func (i *Interactor) PseudoReelo(ctx context.Context, player domain.Player, year int) utils.Error {
	var isParis bool

	//### Steps from 1 to 5
	// There could be more than one category for a year,
	// this could happen in case of namesakes or international results
	categories, err := i.GameRepository.FindCategoriesByYearAndPlayer(ctx, year, player.ID)
	if err != nil {
		i.Logger.Log("PseudoReelo: cannot find categories: %v", err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}

	//var parisIndex int
	for _, c := range categories {

		cities, err := i.ParticipationRepository.FindCitiesByPlayerIDAndGameYearAndCategory(ctx, player.ID, year, c)
		if err != nil {
			i.Logger.Log("PseudoReelo: cannot find cities: %v", err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}

		for _, city := range cities {
			if city == "paris" {
				isParis = true
				break
			}
		}

		// if isParis {
		// 	parisIndex = index
		// 	break
		// }
	}

	//categories = dumbNamesakeGuard(categories, parisIndex, isParis)
	if len(categories) > 0 {
	}

	for _, c := range categories {
		reelo, e := i.oneYearScore(ctx, player, c, year, isParis)
		if !e.IsNil {
			return e
		}
		err := i.ResultRepository.UpdatePseudoReeloByPlayerIDAndGameYearAndCategory(ctx, player.ID, year, c, reelo)
		if err != nil {
			i.Logger.Log("PseudoReelo: failed to update pseudo reelo: %v", err)
			return utils.NewError(err, "E_DB_UPDATE", 500)
		}
	}
	return utils.NewNilError()
}

// Reelo calculates a player's ELO using a custom algorithm
func (i *Interactor) Reelo(ctx context.Context, player domain.Player) (float64, utils.Error) {
	var reelo float64
	var sumOfWeights float64

	// Get some usefull values from db:
	//
	// A list of years in which the player has participated
	// It's used to iterate over the results as well as to check if the
	// anti-exploit mechanism should take effect.
	years, err := i.GameRepository.FindDistinctYearsByPlayerID(ctx, player.ID)
	if err != nil {
		i.Logger.Log("Reelo: cannot find participation years: %v", err)
		return reelo, utils.NewError(err, "E_DB_FIND", 500)
	}
	// The last category known in which the player has participated.
	// It's used to check for category promotion.
	lastKnownCategoryForPlayer, err := i.GameRepository.FindMaxCategoryByPlayerID(ctx, player.ID)
	if err != nil {
		i.Logger.Log("Reelo: cannot find max max category for player %d: %v", player.ID, err)
		return reelo, utils.NewError(err, "E_DB_FIND", 500)
	}
	// The last year known in which the player has participated.
	// It is used to calculate the aging factor and if the anti-exploit mechanism
	// should take effect.
	lastKnownYear, err := i.GameRepository.FindMaxYear(ctx)
	if err != nil {
		i.Logger.Log("Reelo: cannot find max year: %v", err)
		return reelo, utils.NewError(err, "E_DB_FIND", 500)
	}

	// Every year a player has played should have a pseudo-Reelo, which needs
	// to be aged and promoted (if needed) to become a proper Reelo.
	// Finally we want to calculate a weighted average of
	// all those Reeloj and get the final score.
	for _, year := range years {
		// pseudo-Reelo is basically a glorified version of the K*E+D formula.
		pseudoReelo, err := i.ResultRepository.FindPseudoReeloByPlayerIDAndGameYear(ctx, player.ID, year)
		if err != nil {
			i.Logger.Log("Reelo: cannot find max year: %v", err)
			return reelo, utils.NewError(err, "E_DB_FIND", 500)
		}
		// the category for the current year, used to check for category promotion
		cat, err := i.GameRepository.FindCategoryByPlayerIDAndGameYear(ctx, player.ID, year)
		if err != nil {
			i.Logger.Log("Reelo: cannot find category for player %v: %v", player.ID, err)
			return reelo, utils.NewError(err, "E_DB_FIND", 500)
		}
		oldAvg, err := i.ResultRepository.FindAvgPseudoReeloByGameYearAndCategory(context.Background(), year, cat)
		if err != nil {
			i.Logger.Log("Reelo: cannot find old average pseudo reelo: %v", err)
			return reelo, utils.NewError(err, "E_DB_FIND", 500)
		}

		newAvg, err := i.ResultRepository.FindAvgPseudoReeloByGameYearAndCategory(context.Background(), year, lastKnownCategoryForPlayer)
		if err != nil {
			i.Logger.Log("Reelo: cannot find new average pseudo reelo: %v", err)
			return reelo, utils.NewError(err, "E_DB_FIND", 500)
		}
		newMax, e := i.maxPseudoReelo(year, lastKnownCategoryForPlayer)
		if !e.IsNil {
			return reelo, e
		}
		stepSix(&pseudoReelo, lastKnownCategoryForPlayer, cat, year, newAvg, oldAvg, newMax)
		stepSeven(&pseudoReelo, &sumOfWeights, lastKnownYear, year)
		reelo += pseudoReelo
	}

	stepEight(&reelo, sumOfWeights)
	stepNine(&reelo, years, lastKnownYear)
	stepTen(&reelo, years, lastKnownYear)

	return reelo, utils.NewNilError()
}

// oneYearScore is used to calculate a baseScore using steps from 1 to 5.
// This baseScore (a.k.a. pseudo-Reelo) refers to a single year.
// This needs to be calculated for every year the player has played.
func (i *Interactor) oneYearScore(ctx context.Context, player domain.Player, cat string,
	year int, isParis bool) (float64, utils.Error) {
	var baseScore float64

	// the first exercise a player is supposed to solve for the given category
	t, err := i.GameRepository.FindStartByYearAndCategory(ctx, year, cat)
	if err != nil {
		i.Logger.Log("oneYearScore: cannot find starting exercise: %v", err)
		return baseScore, utils.NewError(err, "E_DB_FIND", 500)
	}
	// the last exercise a player is supposed to solve for the given category
	n, err := i.GameRepository.FindEndByYearAndCategory(ctx, year, cat)
	if err != nil {
		i.Logger.Log("oneYearScore: cannot find ending exercise: %v", err)
		return baseScore, utils.NewError(err, "E_DB_FIND", 500)
	}
	// the maximum number of solvable exercises for the given category
	eMax := float64(n - t + 1)
	maxScoreForCat, err := i.ResultRepository.FindMaxScoreByGameYearAndCategory(ctx, year, cat)
	if err != nil {
		i.Logger.Log("oneYearScore: cannot find max score: %v", err)
		return baseScore, utils.NewError(err, "E_DB_FIND", 500)
	}
	// the maximum score obtainable in the given category
	dMax := float64(maxScoreForCat)
	// the player's score for this year-category
	d, err := i.ResultRepository.FindScoreByYearAndPlayerIDAndGameIsParis(ctx, year, player.ID, isParis)
	if err != nil {
		i.Logger.Log("oneYearScore: cannot find score for year %v: %v", year, err)
		return baseScore, utils.NewError(err, "E_DB_FIND", 500)
	}
	// the number of exercises solved by the player for this year-category
	exercises, err := i.ResultRepository.FindExercisesByYearAndPlayerIDAndGameIsParis(ctx, year, player.ID, isParis)
	if err != nil {
		i.Logger.Log("oneYearScore: cannot find exercises for year %v: %v", year, err)
		return baseScore, utils.NewError(err, "E_DB_FIND", 500)
	}

	// This two checks DO NOT solve the problem, it needs manual intervention
	if float64(exercises) > eMax {
		i.Logger.Log("Player %s %s has solved too many exercises (%v > %v) in year %d and category %v\n", player.Name, player.Surname, exercises, eMax, year, cat)
		exercises = 0
	}
	if d > dMax {
		i.Logger.Log("Player %s %s has scored too many  points (%v > %v) in year %d and category %v\n", player.Name, player.Surname, d, dMax, year, cat)
		d = 0
	}
	e := float64(exercises)

	stepOne(&baseScore, e, d)
	stepTwo(&baseScore, isParis)
	stepThree(&baseScore, t, n, d, e, eMax, dMax)
	avgCatScore, err := i.ResultRepository.FindAvgScoreByGameYear(context.Background(), year, int(exercisesCostant))
	if err != nil {
		i.Logger.Log("oneYearScore: cannot find average score: %v", err)
		return baseScore, utils.NewError(err, "E_DB_FIND", 500)
	}
	stepFour(&baseScore, year, avgCatScore)
	stepFive(&baseScore)

	return baseScore, utils.NewNilError()
}

func (i *Interactor) maxPseudoReelo(year int, cat string) (float64, utils.Error) {
	var pseudoReelo float64

	t, err := i.GameRepository.FindStartByYearAndCategory(context.Background(), year, cat)
	if err != nil {
		i.Logger.Log("maxPseudoReelo: cannot find starting exercise for year/category (%d/%s): %v", year, cat, err)
		return pseudoReelo, utils.NewError(err, "E_DB_FIND", 500)
	}
	n, err := i.GameRepository.FindEndByYearAndCategory(context.Background(), year, cat)
	if err != nil {
		i.Logger.Log("maxPseudoReelo: cannot find ending exercise: %v", err)
		return pseudoReelo, utils.NewError(err, "E_DB_FIND", 500)
	}

	// Here eMax correspond to e
	eMax := float64(n - t + 1)
	e := eMax
	// Here dMax correspond to d
	maxScoreForCat, err := i.ResultRepository.FindMaxScoreByGameYearAndCategory(context.Background(), year, cat)
	if err != nil {
		i.Logger.Log("maxPseudoReelo: cannot find max score: %v", err)
		return pseudoReelo, utils.NewError(err, "E_DB_FIND", 500)
	}
	// the maximum score obtainable in the given category
	dMax := float64(maxScoreForCat)
	d := dMax

	stepOne(&pseudoReelo, e, d)
	stepTwo(&pseudoReelo, true)
	stepThree(&pseudoReelo, t, n, d, e, eMax, dMax)
	avgCatScore, err := i.ResultRepository.FindAvgScoreByGameYear(context.Background(), year, int(exercisesCostant))
	if err != nil {
		i.Logger.Log("maxPseudoReelo: cannot find average score: %v", err)
		return pseudoReelo, utils.NewError(err, "E_DB_FIND", 500)
	}
	stepFour(&pseudoReelo, year, avgCatScore)
	stepFive(&pseudoReelo)

	return pseudoReelo, utils.NewNilError()
}

func contains(array []int, item int) bool {
	for _, e := range array {
		if e == item {
			return true
		}
	}
	return false
}

func dumbNamesakeGuard(categories []string, parisIndex int, isParis bool) []string {
	if (len(categories) > 1 && !isParis) || (len(categories) > 2 && isParis) {
		maxC := category.FromString("CE")
		for _, c := range categories {
			if tmp := category.FromString(c); tmp >= maxC {
				maxC = tmp
			}
		}
		log.Printf("FOUND NAMESAKE - Categories: %v\nConsidering only the highest result: %v.\n", categories, maxC)
		if isParis {
			return []string{categories[parisIndex], maxC.String()}
		}
		return []string{maxC.String()}
	}
	return categories
}
