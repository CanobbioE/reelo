package repository

import "context"

// PARTECIPATIONREPO is the handler name
const PARTECIPATIONREPO = "PartecipationRepo"

// DbPartecipationRepo id the repository for Partecipations
type DbPartecipationRepo DbRepo

// NewDbPartecipationRepo istanciates and returns a Partecipation repository
func NewDbPartecipationRepo(dbHandlers map[string]DbHandler) *DbPartecipationRepo {
	return &DbPartecipationRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[PARTECIPATIONREPO],
	}
}

// FindCitiesByPlayerIDAndGameYearAndCategory returns a list of cities for
// the given player's ID, game's year and game's category
func (db *DbPartecipationRepo) FindCitiesByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) ([]string, error) {
}

// FindAll retrieves all the Partecipations in the repository, paginating the results
func (db *DbPartecipationRepo) FindAll(ctx context.Context, page, size int) ([]Partecipation, error) {
}
