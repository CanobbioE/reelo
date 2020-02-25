package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/CanobbioE/reelo/backend/domain"
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

// CalculateAllReelo recalculates the Reelo score
// for every single player in the database.
// If the doPseudo is true, then the pseudo-reelo gets recalculated aswell
func (i *Interactor) CalculateAllReelo(doPseudo bool) utils.Error {
	start := time.Now()
	errs, ctx := errgroup.WithContext(context.Background())
	if doPseudo {
		i.Logger.Log("(Re)calculating pseudo-Reelo...")
	}
	i.Logger.Log("(Re)calculating Reelo...")

	i.InitCostants()
	ids, err := i.PlayerRepository.FindAllIDs(ctx)
	if err != nil {
		i.Logger.Log("PlayersCount: cannot find players count: %v", err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}

	for _, id := range ids {
		player, err := i.PlayerRepository.FindByID(ctx, id)
		if err != nil {
			i.Logger.Log("PlayersCount: cannot find players count: %v", err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}
		if player.Name == "" || player.Surname == "" {
			continue
		}

		errs.Go(func() error {
			err := i.CalculatePlayerReelo(player, doPseudo)
			if !err.IsNil {
				return fmt.Errorf(err.Message)
			}
			return nil
		})
	}

	err = errs.Wait()
	end := time.Now()
	i.Logger.Log("recalculating reelo for %v players took %v", len(ids), end.Sub(start))
	return utils.NewError(err, "E_GENERIC", 500)

}

// CalculatePlayerReelo recalculates the reelo for a single user
func (i *Interactor) CalculatePlayerReelo(player domain.Player, doPseudo bool) utils.Error {
	ctx := context.Background()

	if doPseudo {
		years, err := i.GameRepository.FindDistinctYearsByPlayerID(ctx, player.ID)
		if err != nil {
			i.Logger.Log("CalculatePlayerReelo: cannot find years for %v: %v", player.ID, err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}

		for _, year := range years {
			err := i.PseudoReelo(ctx, player, year)
			if !err.IsNil {
				return err
			}
		}
	}

	elo, e := i.Reelo(ctx, player)
	if !e.IsNil {
		return e
	}
	player.Reelo = elo
	err := i.PlayerRepository.UpdateReelo(ctx, player)
	if err != nil {
		i.Logger.Log("Error updating reelo: %v", err)
		return utils.NewError(err, "E_DB_UPDATE", 500)
	}
	return utils.NewNilError()
}
