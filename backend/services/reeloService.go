package services

import (
	"context"
	"log"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/services/elo"
)

// CalculateAllReelo recalculates the Reelo score
// for every single player in the database
func CalculateAllReelo() error {
	log.Println("Recalculating reelo")
	elo.InitCostants()
	ctx := context.Background()
	db := rdb.NewDB()
	players, err := db.AllPlayers(ctx)
	if err != nil {
		return err
	}
	for _, player := range players {
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
	return nil
}
