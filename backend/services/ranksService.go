package services

import (
	"context"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/dto"
)

// GetRanks returns a list of all the ranks in the database
func GetRanks(page, size int) ([]dto.Rank, error) {
	db := rdb.Instance()
	return db.AllRanks(context.Background(), page, size)
}

// GetPlayersCount returns how many players are stored in the DB
func GetPlayersCount() (int, error) {
	db := rdb.Instance()
	return db.CountAllPlayers(context.Background())
}

// GetHistory returns the history details for the players
func GetHistory(name, surname string) (dto.PlayerHistory, error) {
	db := rdb.Instance()
	return db.PlayerHistory(context.Background(), name, surname)
}
