package usecases

import (
	"context"
)

// SlimPartecipation represents a simplified partecipation relationship
type SlimPartecipation struct {
	City         string  `json:"city"`
	Category     string  `json:"category"`
	IsParis      bool    `json:"isParis"`
	Year         int     `json:"year"`
	MaxExercises int     `json:"eMax"`
	MaxScore     int     `json:"dMax"`
	Score        int     `json:"d"`
	Exercises    int     `json:"e"`
	Time         int     `json:"time"`
	Position     int     `json:"position"`
	PseudoReelo  float64 `json:"pseudoReelo"`
}

// History is a collection of simplified partecipations
type History []SlimPartecipation

// SlimPartecipationByYear is a slim partecipation indexed by parteciaption year to simplify resarch
type SlimPartecipationByYear map[int]SlimPartecipation

// HistoryByYear is an history indexed by parteciaption year to simplify resarch
type HistoryByYear map[int]History

// HistoryRepository is the interface for the persistency container
type HistoryRepository interface {
	FindByPlayerIDAndYear(ctx context.Context, id, y int) (SlimPartecipationByYear, error)
	FindByPlayerIDOrderByYear(ctx context.Context, id int) (HistoryByYear, []int, error)
}

// HistorySwitcheroo(ctx context.Context, oldID, newID int, newHistory []History) error

// IsEqual compares two slimPartecipations and returns true if they are equivalent,
// false otherwise
func (s *SlimPartecipation) IsEqual(b SlimPartecipation) bool {
	return s.Year == b.Year && s.Category == b.Category && s.City == b.City && s.IsParis == b.IsParis
}
