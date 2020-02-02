package repository

import (
	"context"
)

// HISTORYREPO is the handler name
const HISTORYREPO = "historyRepo"

// DbHistoryRepo id the repository for Historys
type DbHistoryRepo DbRepo

// NewDbHistoryRepo istanciates and returns a History repository
func NewDbHistoryRepo(dbHandlers map[string]DbHandler) *DbHistoryRepo {
	return &DbHistoryRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[HISTORYREPO],
	}
}

// FindByPlayerID retrieves the history for the given player id
func (db *DbCostantsRepo) FindByPlayerID(ctx context.Context, id int) (History, []int, error) {}
