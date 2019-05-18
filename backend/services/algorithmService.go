package services

import (
	"context"

	"github.com/CanobbioE/reelo/backend/api"
	rdb "github.com/CanobbioE/reelo/backend/db"
)

// UpdateAlgorithm updates some varaibles used for the reelo algorithm
func UpdateAlgorithm(ctx context.Context, c api.Costants) error {
	db := rdb.NewDB()
	defer db.Close()

	return db.UpdateCostants(ctx, rdb.Costants(c))

}

// GetCostants fetch the current values for the variables used
// in the reelo algorithm
func GetCostants() (api.Costants, error) {
	db := rdb.NewDB()
	defer db.Close()
	c, err := db.ReeloCostants()
	if err != nil {
		return api.Costants(c), err
	}

	return api.Costants(c), nil

}
