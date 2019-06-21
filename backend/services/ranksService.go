package services

import (
	"context"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/dto"
)

// GetRanks returns a list of all the ranks in the database
func GetRanks() ([]dto.Rank, error) {
	db := rdb.NewDB()
	defer db.Close()
	return db.AllRanks(context.Background())
}
