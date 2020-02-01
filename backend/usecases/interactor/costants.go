package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
)

// UpdateCostants updates some varaibles used to calculate the players' reelo
func (i *Interactor) UpdateCostants(costants domain.Costants) error {
	return i.CostantsRepository.UpdateAll(context.Background(), costants)
}

// ListCostants returns the current values for the variables used to calculate
// the players' reelo
func (i *Interactor) ListCostants() (domain.Costants, error) {
	return i.CostantsRepository.FindAll(context.Background())
}
