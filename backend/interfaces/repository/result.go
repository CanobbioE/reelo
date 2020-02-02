package repository

import (
	"context"
)

// ResultREPO is the handler name
const ResultREPO = "ResultRepo"

// DbResultRepo id the repository for Results
type DbResultRepo DbRepo

// NewDbResultRepo istanciates and returns a Result repository
func NewDbResultRepo(dbHandlers map[string]DbHandler) *DbResultRepo {
	return &DbResultRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[ResultREPO],
	}
}

// FindAllByPlayerID does stuff
func (db *DbResultRepo) FindAllByPlayerID(ctx context.Context, id int) ([]Result, error)

// FindExercisesByPlayerID does stuff
func (db *DbResultRepo) FindExercisesByPlayerID(ctx context.Context, id int) (int, error)

// FindScoresByYearAndPlayerIDAndGameIsParis does stuff
func (db *DbResultRepo) FindScoresByYearAndPlayerIDAndGameIsParis(ctx context.Context, y, id int, ip bool) (float64, error)

// FindExercisesByYearAndPlayerIDAndGameIsParis does stuff
func (db *DbResultRepo) FindExercisesByYearAndPlayerIDAndGameIsParis(ctx context.Context, y, id int, ip bool) (int, error)

// FindAvgScoreByGameYear does stuff
func (db *DbResultRepo) FindAvgScoreByGameYear(ctx context.Context, y, k int) (float64, error)

// FindAvgPseudoReeloByGameYearAndCategory does stuff
func (db *DbResultRepo) FindAvgPseudoReeloByGameYearAndCategory(ctx context.Context, y int, c string) (float64, error)

// FindMaxScoreByGameYearAndCategory does stuff
func (db *DbResultRepo) FindMaxScoreByGameYearAndCategory(ctx context.Context, y int, c string) (int, error)

// FindPseudoReeloByPlayerIDAndGameYear does stuff
func (db *DbResultRepo) FindPseudoReeloByPlayerIDAndGameYear(ctx context.Context, id, y int) (float64, error)

// FindIDByPlayerIDAndGameYearAndCategory does stuff
func (db *DbResultRepo) FindIDByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) (int, error)

// UpdatePseudoReeloByPlayerIDAndGameYearAndCategory does stuff
func (db *DbResultRepo) UpdatePseudoReeloByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string, pr float64) error

// DeleteByGameID does stuff
func (db *DbResultRepo) DeleteByGameID(ctx context.Context, id int) error
