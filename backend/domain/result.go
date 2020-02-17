package domain

import (
	"context"
)

// Result represents the result relation as it is stored in the db
type Result struct {
	ID          int     `json:"id"`
	Exercises   int     `json:"exercises"`
	Time        int     `json:"time"`
	Score       int     `json:"score"`
	Position    int     `json:"position"`
	PseudoReelo float64 `json:"pseudoReelo"`
}

// ResultRepository is the interface for the persistency container
type ResultRepository interface {
	Store(ctx context.Context, r Result) (int64, error)
	FindAllByPlayerID(ctx context.Context, id int) ([]Result, error)
	FindExercisesByPlayerID(ctx context.Context, id int) (int, error)
	FindScoreByYearAndPlayerIDAndGameIsParis(ctx context.Context, y, id int, ip bool) (float64, error)
	FindExercisesByYearAndPlayerIDAndGameIsParis(ctx context.Context, y, id int, ip bool) (int, error)
	FindAvgScoreByGameYear(ctx context.Context, y, k int) (float64, error)
	FindAvgPseudoReeloByGameYearAndCategory(ctx context.Context, y int, c string) (float64, error)
	FindMaxScoreByGameYearAndCategory(ctx context.Context, y int, c string) (float64, error)
	FindPseudoReeloByPlayerIDAndGameYear(ctx context.Context, id, y int) (float64, error)
	FindIDByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) (int, error)
	UpdatePseudoReeloByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string, pr float64) error
	DeleteByGameID(ctx context.Context, id int) error
	FindByPlayerIDAndGameYear(ctx context.Context, id, y int) (Result, error)
}
