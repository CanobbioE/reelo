package services

import (
	"context"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/dto"
)

// UpdateAlgorithm updates some varaibles used for the reelo algorithm
func UpdateAlgorithm(ctx context.Context, c dto.Costants) error {
	db := rdb.NewDB()
	defer db.Close()

	return db.UpdateCostants(ctx, rdb.Costants(c))
}

// GetCostants fetch the current values for the variables used
// in the reelo algorithm
func GetCostants() (dto.Costants, error) {
	db := rdb.NewDB()
	defer db.Close()
	c, err := db.ReeloCostants()
	if err != nil {
		return dto.Costants(c), err
	}

	return dto.Costants(c), nil

}
