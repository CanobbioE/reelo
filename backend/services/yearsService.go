package services

import (
	"context"

	rdb "github.com/CanobbioE/reelo/backend/db"
)

// GetYears returns a list of all the stored years
func GetYears() ([]int, error) {
	db := rdb.NewDB()
	defer db.Close()
	return db.AllYears(context.Background())
}
