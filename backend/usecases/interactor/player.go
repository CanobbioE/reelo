package interactor

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/pkg/reelo"
	"github.com/CanobbioE/reelo/backend/utils"
	"golang.org/x/sync/errgroup"
)

// PlayersCount returns how many players are stored in the DB
func (i *Interactor) PlayersCount() (int, utils.Error) {
	count, err := i.PlayerRepository.FindCountAll(context.Background())
	if err != nil {
		i.Logger.Log("PlayersCount: cannot find players count: %v", err)
		return count, utils.NewError(err, "E_DB_FIND", 500)
	}
	return count, utils.NewNilError()
}

// CalculateAllReelo (re)calculates the Reelo score for every player in the database.
func (i *Interactor) CalculateAllReelo() utils.Error {
	// Initialize Constants
	c := reelo.NewConstants()
	constFromDB, err := i.CostantsRepository.FindAll(context.Background())
	if err != nil {
		i.Logger.Log("Error initializing constants: %v", err)
		i.Logger.Log("Falling back to the hardcoded configuration\n")
	} else {
		c.StartingYear = constFromDB.StartingYear
		c.ExercisesCostant = constFromDB.ExercisesCostant
		c.PFinal = constFromDB.PFinal
		c.MultiplicativeFactor = constFromDB.MultiplicativeFactor
		c.AntiExploit = constFromDB.AntiExploit
		c.NoParticipationPenalty = constFromDB.NoParticipationPenalty
	}
	reelo.InitConstants(c)

	// Calculate Pseudo reelo for all players. Must do this before calculating
	// the players Reelo, since it assumes all players have already a pseudo reelo.
	err1 := execFuncOnAllPlayers(playerPseudoReelo, i)
	if !err1.IsNil {
		return err1
	}

	// Calculate Reelo for all players
	err1 = execFuncOnAllPlayers(playerReelo, i)
	if !err1.IsNil {
		return err1
	}
	return utils.NewNilError()
}

func execFuncOnAllPlayers(fx func(i *Interactor, p domain.Player) utils.Error, i *Interactor) utils.Error {
	var g errgroup.Group
	ctx := context.Background()
	// Recover all players IDs from DB
	ids, err := i.PlayerRepository.FindAllIDs(ctx)
	if err != nil {
		i.Logger.Log("CalculateAllReelo: cannot find all players ids: %v", err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}

	for _, id := range ids {
		player, err := i.PlayerRepository.FindByID(ctx, id)
		if err != nil {
			i.Logger.Log("CalculateAllReelo: cannot find player by id %d: %v", id, err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}

		if player.Name == "" || player.Surname == "" {
			continue
		}

		g.Go(func() error {
			err := fx(i, player)
			if !err.IsNil {
				return fmt.Errorf(err.Message)
			}
			return nil
		})
	}

	err = g.Wait()
	return utils.NewError(err, "E_GENERIC", 500)
}

// playerReelo recalculates the reelo for a single user
func playerReelo(i *Interactor, player domain.Player) utils.Error {
	var args reelo.Args
	var err error
	args.History = make(map[int]reelo.Details)
	ctx := context.Background()

	// Find all years the player has participated
	args.Years, err = i.GameRepository.FindDistinctYearsByPlayerID(ctx, player.ID)
	if err != nil {
		i.Logger.Log("CalculatePlayerReelo: cannot find years for %v: %v", player.ID, err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}

	// Update the Reelo using the newly (re)calculated one
	player.Reelo = reelo.Calculate(args)
	err = i.PlayerRepository.UpdateReelo(ctx, player)
	if err != nil {
		i.Logger.Log("CalculatePlayerReelo: cannot update reelo: %v", err)
		return utils.NewError(err, "E_DB_UPDATE", 500)
	}
	return utils.NewNilError()
}

// playerPseudoReelo recalculates the pseudo reelo for a single user
func playerPseudoReelo(i *Interactor, player domain.Player) utils.Error {
	ctx := context.Background()
	// Find all years the player has participated
	years, err := i.GameRepository.FindDistinctYearsByPlayerID(ctx, player.ID)
	if err != nil {
		i.Logger.Log("CalculatePlayerReelo: cannot find years for %v: %v", player.ID, err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}

	// Iterate over the years, each pair year-category should have a pseudo reelo
	for _, year := range years {
		var pseudoArgs reelo.SlimArgs

		// Find all categories for the given year
		categories, err := i.GameRepository.FindCategoriesByYearAndPlayer(ctx, year, player.ID)
		if err != nil {
			i.Logger.Log("PseudoReelo: cannot find categories: %v", err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}

		// We might have more than one category per year, which means the
		// result might be from an international result.
		for _, cat := range categories {

			// Determine if it is an international result by checking the city
			cities, err := i.ParticipationRepository.FindCitiesByPlayerIDAndGameYearAndCategory(ctx, player.ID, year, cat)
			if err != nil {
				i.Logger.Log("cannot find cities: %v", err)
				return utils.NewError(err, "E_DB_FIND", 500)
			}
			for _, city := range cities {
				if city == "paris" {
					pseudoArgs.IsParis = true
					break
				}

				// Find the first exercise to be solved for the given category
				pseudoArgs.Start, err = i.GameRepository.FindStartByYearAndCategory(ctx, year, cat)
				if err != nil {
					i.Logger.Log("cannot find starting exercise: %v", err)
					return utils.NewError(err, "E_DB_FIND", 500)
				}

				// Find the last exercise to be solved for the given category
				pseudoArgs.End, err = i.GameRepository.FindEndByYearAndCategory(ctx, year, cat)
				if err != nil {
					i.Logger.Log("cannot find ending exercise: %v", err)
					return utils.NewError(err, "E_DB_FIND", 500)
				}

				// Find the maximum score obtainable for the given category
				pseudoArgs.MaxScoreForCategory, err = i.ResultRepository.FindMaxScoreByGameYearAndCategory(ctx, year, cat)
				if err != nil {
					i.Logger.Log("cannot find max score: %v", err)
					return utils.NewError(err, "E_DB_FIND", 500)
				}

				// Find the player's score
				pseudoArgs.Score, err = i.ResultRepository.FindScoreByYearAndPlayerIDAndGameIsParis(ctx, year, player.ID, pseudoArgs.IsParis)
				if err != nil {
					i.Logger.Log("cannot find score for year %v: %v", year, err)
					return utils.NewError(err, "E_DB_FIND", 500)
				}
				// Find the number of solved exercises
				pseudoArgs.Exercises, err = i.ResultRepository.FindExercisesByYearAndPlayerIDAndGameIsParis(ctx, year, player.ID, pseudoArgs.IsParis)
				if err != nil {
					i.Logger.Log("cannot find exercises for year %v: %v", year, err)
					return utils.NewError(err, "E_DB_FIND", 500)
				}

				pseudoReelo := reelo.CalculatePseudo(pseudoArgs)
				// Update the pseudoReelo using the newly (re)calculated one
				err = i.ResultRepository.UpdatePseudoReeloByPlayerIDAndGameYearAndCategory(ctx, player.ID, year, cat, pseudoReelo)
				if err != nil {
					i.Logger.Log("PseudoReelo: failed to update pseudo reelo: %v", err)
					return utils.NewError(err, "E_DB_UPDATE", 500)
				}
			}
		}
	}
	return utils.NewNilError()
}
