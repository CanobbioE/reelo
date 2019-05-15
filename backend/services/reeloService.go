package services

import (
	"context"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/services/elo"
)

// CalculateAllReelo recalculates the Reelo score
// for every single player in the database
func CalculateAllReelo() {
	ctx := context.Background()
	db := rdb.NewDB()
	players := db.GetAllPlayers(ctx)
	for _, player := range players {
		player.Reelo = int(elo.Reelo(player.Name, player.Surname))
		db.UpdateReelo(ctx, player)
	}
}
