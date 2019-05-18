package services

import (
	"context"
	"log"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/services/elo"
)

// CalculateAllReelo recalculates the Reelo score
// for every single player in the database
func CalculateAllReelo() {
	elo.InitCostants()
	ctx := context.Background()
	db := rdb.NewDB()
	players := db.AllPlayers(ctx)
	for _, player := range players {
		player.Reelo = int(elo.Reelo(ctx, player.Name, player.Surname))
		err := db.UpdateReelo(ctx, player)
		if err != nil {
			log.Printf("Error updating reelo: %v", err)
		}
	}
}
