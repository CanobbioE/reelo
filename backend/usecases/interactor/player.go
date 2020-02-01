package interactor

import (
	"context"
	"time"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases/elo"
	"golang.org/x/sync/errgroup"
)

// PlayersCount returns how many players are stored in the DB
func (i *Interactor) PlayersCount() (int, error) {
	return i.PlayerRepository.FindCountAll(context.Background())
}

// CalculateAllReelo recalculates the Reelo score
// for every single player in the database.
// If the doPseudo is true, then the pseudo-reelo gets recalculated aswell
func (i *Interactor) CalculateAllReelo(doPseudo bool) error {
	start := time.Now()
	errs, ctx := errgroup.WithContext(context.Background())
	if doPseudo {
		i.Logger.Log("(Re)calculating pseudo-Reelo...")
	}
	i.Logger.Log("(Re)calculating Reelo...")

	elo.InitCostants()
	ids, err := i.PlayerRepository.FindAllIDs(ctx)
	if err != nil {
		return err
	}

	for _, id := range ids {
		player, err := i.PlayerRepository.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if player.Name == "" || player.Surname == "" {
			continue
		}

		errs.Go(func() error { return i.CalculatePlayerReelo(player, doPseudo) })
	}

	err = errs.Wait()
	end := time.Now()
	i.Logger.Log("recalculating reelo for %v players took %v", len(ids), end.Sub(start))
	return err

}

// CalculatePlayerReelo recalculates the reelo for a single user
func (i *Interactor) CalculatePlayerReelo(player domain.Player, doPseudo bool) error {
	ctx := context.Background()

	if doPseudo {
		years, err := i.GameRepository.FindDistinctYearsByPlayerID(ctx, player.ID)
		if err != nil {
			i.Logger.Log(player.Name, player.Surname, err)
			return err
		}

		for _, year := range years {
			err := elo.PseudoReelo(ctx, player.Name, player.Surname, year)
			if err != nil {
				return err
			}
		}
	}

	elo, err := elo.Reelo(ctx, player.Name, player.Surname)
	if err != nil {
		return err
	}
	player.Reelo = elo
	err = i.PlayerRepository.UpdateReelo(ctx, player)
	if err != nil {
		i.Logger.Log("Error updating reelo: %v", err)
		return err
	}
	return nil
}
