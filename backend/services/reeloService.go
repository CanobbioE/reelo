package services

import (
	"context"
	"log"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/services/elo"
)

// CalculateAllReelo recalculates the Reelo score
// for every single player in the database.
// If the doPseudo is true, then the pseudo-reelo gets recalculated aswell
func CalculateAllReelo(doPseudo bool) error {
	if doPseudo {
		log.Println("(Re)calculating pseudo-Reelo...")
	}
	log.Println("(Re)calculating Reelo...")

	elo.InitCostants()
	ctx := context.Background()
	db := rdb.NewDB()
	players, err := db.AllPlayers(ctx)
	if err != nil {
		return err
	}

	for _, player := range players {
		if player.Name == "" || player.Surname == "" {
			continue
		}
		if doPseudo {
			years, err := db.PlayerPartecipationYears(ctx, player.Name, player.Surname)
			if err != nil {
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
	}
	log.Println("Reelo (re)calculated.")
	return nil
}
