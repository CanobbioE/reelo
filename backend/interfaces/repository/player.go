package repository

import (
	"context"
)

// PLAYERREPO is the handler name
const PLAYERREPO = "playerRepo"

// DbPlayerRepo id the repository for Players
type DbPlayerRepo DbRepo

// NewDbPlayerRepo istanciates and returns a Player repository
func NewDbPlayerRepo(dbHandlers map[string]DbHandler) *DbPlayerRepo {
	return &DbPlayerRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[PLAYERREPO],
	}
}

// FindIDByNameAndSurname does stuff
func (db *DbPlayerRepo) FindIDByNameAndSurname(ctx context.Context, n, s string) (int, error) {}

// FindAllIDs does stuff
func (db *DbPlayerRepo) FindAllIDs(ctx context.Context) ([]int, error) {}

// FindByID does stuff
func (db *DbPlayerRepo) FindByID(ctx context.Context, id int) (Player, error) {}

// FindAll does stuff
func (db *DbPlayerRepo) FindAll(ctx context.Context, page, size int) ([]Player, error) {}

// FindCountAll does stuff
func (db *DbPlayerRepo) FindCountAll(ctx context.Context) (int, error) {}

// UpdateReelo does stuff
func (db *DbPlayerRepo) UpdateReelo(ctx context.Context, p Player) error {}

// UpdateAccent does stuff
func (db *DbPlayerRepo) UpdateAccent(ctx context.Context, a string) {}

// CheckExistenceByID does stuff
func (db *DbPlayerRepo) CheckExistenceByID(ctx context.Context, id int) bool {}
