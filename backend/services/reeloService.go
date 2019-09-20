package services

import (
	"context"
	"log"
	"time"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/services/elo"
	"golang.org/x/sync/errgroup"
)

// CalculateAllReelo recalculates the Reelo score
// for every single player in the database.
// If the doPseudo is true, then the pseudo-reelo gets recalculated aswell
func CalculateAllReelo(doPseudo bool) error {
	start := time.Now()
	errs, ctx := errgroup.WithContext(context.Background())
	if doPseudo {
		log.Println("(Re)calculating pseudo-Reelo...")
	}
	log.Println("(Re)calculating Reelo...")

	elo.InitCostants()
	db := rdb.Instance()
	ids, err := db.AllPlayersID(ctx)
	if err != nil {
		return err
	}

	for _, id := range ids {
		player, err := db.Player(ctx, id)
		if err != nil {
			return err
		}
		if player.Name == "" || player.Surname == "" {
			continue
		}

		errs.Go(func() error { return playerReelo(player, doPseudo) })
	}

	err = errs.Wait()
	end := time.Now()
	log.Printf("recalculating reelo for %v players took %v", len(ids), end.Sub(start))
	return err
}

// PlayerReelo recalculates the reelo for a single user
func playerReelo(player rdb.Player, doPseudo bool) error {
	db := rdb.Instance()
	ctx := context.Background()

	if doPseudo {
		years, err := db.PlayerPartecipationYears(ctx, player.Name, player.Surname)
		if err != nil {
			log.Println(player.Name, player.Surname, err)
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
	player.Reelo = int(elo)
	err = db.UpdateReelo(ctx, player)
	if err != nil {
		log.Printf("Error updating reelo: %v", err)
		return err
	}
	return nil
}
