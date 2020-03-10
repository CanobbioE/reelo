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
// If the doPseudo is true, then the pseudo-reelo gets recalculated as well
func (i *Interactor) CalculateAllReelo(doPseudo bool) utils.Error {
	start := time.Now()
	errs, ctx := errgroup.WithContext(context.Background())

	i.InitCostants()
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
	i.Logger.Log("CalculateAllReelo: recalculating reelo for %v players took %v", len(ids), end.Sub(start))
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
		i.Logger.Log("CalculatePlayerReelo: cannot update reelo: %v", err)
		return utils.NewError(err, "E_DB_UPDATE", 500)
	}
	return utils.NewNilError()
}

// PlayersCleanUp removes all the players that don't have any
// stored participation.
func (i *Interactor) PlayersCleanUp() utils.Error {
	ids, err := i.PlayerRepository.FindAllIDsWhereIDNotInParticipation(context.Background())
	if err != nil {
		i.Logger.Log("PlayersCleanUp: cannot find ids: %v", err)
		return utils.NewError(err, "E_DB_FIND", 500)
	}

	for _, id := range ids {
		err := i.PlayerRepository.DeleteByID(context.Background(), id)
		if err != nil {
			i.Logger.Log("PlayersCleanUp: cannot delete player %d: %v", id, err)
			return utils.NewError(err, "E_DB_DELETE", 500)
		}
	}
	return utils.NewNilError()
}
