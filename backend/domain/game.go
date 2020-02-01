package domain

import "context"

// Game represents the game entity as it is stored in the db
type Game struct {
	ID       int    `json:"id"`
	Year     int    `json:"year"`
	Category string `json:"category"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
	IsParis  bool   `json:"isParis"`
}

// GameRepository is the interface for the persistency container
type GameRepository interface {
	FindIDByYearAndCategoryAndIsParis(ctx context.Context, y int, c string, ip bool) (int, error)
	FindDistinctYearsByPlayerID(ctx context.Context, id int) ([]int, error)
	FindCategoriesByYearAndPlayer(ctx context.Context, y, id int) ([]string, error)
	FindMaxYear(ctx context.Context) (int, error)
	FindMaxCategoryByPlayerID(ctx context.Context, id int) (string, error)
	FindStartByYearAndCategory(ctx context.Context, y int, c string) (int, error)
	FindEndByYearAndCategory(ctx context.Context, y int, c string) (int, error)
	FindCategoryByPlayerIDAndGameYear(ctx context.Context, id, y int) (string, error)
	FindAllYears(ctx context.Context) ([]int, error)
}
