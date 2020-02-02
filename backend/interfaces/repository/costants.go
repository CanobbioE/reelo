package repository

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
)

// COSTANTSREPO is the handler name
const COSTANTSREPO = "costantsRepo"

// DbCostantsRepo id the repository for Costantss
type DbCostantsRepo DbRepo

// NewDbCostantsRepo istanciates and returns a Costants repository
func NewDbCostantsRepo(dbHandlers map[string]DbHandler) *DbCostantsRepo {
	return &DbCostantsRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[CostantsREPO],
	}
}

// FindAll retrieves all the costants in the repository.
// This is confusing due to having only one entry in the "costants" table
func (db *DbCostantsRepo) FindAll(ctx context.Context) (domain.Costants, error) {
}

// UpdateAll updates all the costants in the repository
func (db *DbCostantsRepo) UpdateAll(ctx context.Context, c domain.Costants) error {
}
