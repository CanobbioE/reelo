package usecases

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
)

// ResultDetails represents details on a parteciaption
type ResultDetails struct {
	Partecipation domain.Partecipation `json:"partecipation"`
	MaxExercises  int                  `json:"eMax"`
	MaxScore      int                  `json:"dMax"`
}

// History is a map of results details indexed by year
type History map[int]ResultDetails

// HistoryRepository is the interface for the persistency container
type HistoryRepository interface {
	FindByPlayerID(ctx context.Context, id int) (History, []int, error)
}

// HistorySwitcheroo(ctx context.Context, oldID, newID int, newHistory []History) error
