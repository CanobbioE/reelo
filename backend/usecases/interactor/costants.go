package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
)

// UpdateCostants updates some varaibles used to calculate the players' reelo
func (i *Interactor) UpdateCostants(costants domain.Costants) utils.Error {
	err := i.CostantsRepository.UpdateAll(context.Background(), costants)
	if err != nil {
		i.Logger.Log("UpdateCostants: cannot update all: %v", err)
		return utils.NewError(err, "E_DB_UPDATE", 500)
	}
	return utils.NewNilError()
}

// ListCostants returns the current values for the variables used to calculate
// the players' reelo
func (i *Interactor) ListCostants() (domain.Costants, utils.Error) {
	costants, err := i.CostantsRepository.FindAll(context.Background())
	if err != nil {
		i.Logger.Log("ListCostants: cannot find all: %v", err)
		return costants, utils.NewError(err, "E_DB_FIND", 500)
	}
	return costants, utils.NewNilError()
}
