package repository

import "context"

// GAMEREPO is the handler name
const GAMEREPO = "gameRepo"

// DbGameRepo id the repository for Games
type DbGameRepo DbRepo

// NewDbGameRepo istanciates and returns a Game repository
func NewDbGameRepo(dbHandlers map[string]DbHandler) *DbGameRepo {
	return &DbGameRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[GAMEREPO],
	}
}

// FindIDByYearAndCategoryAndIsParis does stuff
func (db *DbGameRepo) FindIDByYearAndCategoryAndIsParis(ctx context.Context, y int, c string, ip bool) (int, error) {
}

// FindDistinctYearsByPlayerID does stuff
func (db *DbGameRepo) FindDistinctYearsByPlayerID(ctx context.Context, id int) ([]int, error) {}

// FindCategoriesByYearAndPlayer does stuff
func (db *DbGameRepo) FindCategoriesByYearAndPlayer(ctx context.Context, y, id int) ([]string, error) {
}

// FindMaxYear does stuff
func (db *DbGameRepo) FindMaxYear(ctx context.Context) (int, error) {}

// FindMaxCategoryByPlayerID does stuff
func (db *DbGameRepo) FindMaxCategoryByPlayerID(ctx context.Context, id int) (string, error) {}

// FindStartByYearAndCategory does stuff
func (db *DbGameRepo) FindStartByYearAndCategory(ctx context.Context, y int, c string) (int, error) {}

// FindEndByYearAndCategory does stuff
func (db *DbGameRepo) FindEndByYearAndCategory(ctx context.Context, y int, c string) (int, error) {}

// FindCategoryByPlayerIDAndGameYear does stuff
func (db *DbGameRepo) FindCategoryByPlayerIDAndGameYear(ctx context.Context, id, y int) (string, error) {
}

// FindAllYears does stuff
func (db *DbGameRepo) FindAllYears(ctx context.Context) ([]int, error) {}
