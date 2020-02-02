package usecases

import "context"

// History is a map of results details indexed by year
type History map[int]struct {
	Partecipation Partecipation `json:"partecipation"`
	MaxExercises  int           `json:"eMax"`
	MaxScore      int           `json:"dMax"`
}

// HistoryRepository is the interface for the persistency container
type HistoryRepository interface {
	FindByPlayerID(ctx context.Context, id int) (History, []int, error)
}

// HistorySwitcheroo(ctx context.Context, oldID, newID int, newHistory []History) error
