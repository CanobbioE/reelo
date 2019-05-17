package services

import (
	"context"

	"github.com/CanobbioE/reelo/backend/api"
	rdb "github.com/CanobbioE/reelo/backend/db"
)

// GetRanks returns a list of all the ranks in the database
func GetRanks() ([]api.Rank, error) {
	db := rdb.NewDB()
	defer db.Close()
	return db.AllRanks(context.Background())
}
