package usecases

import (
	"context"
)

// TODO: soon obsolete

// SlimParticipation represents a simplified participation relationship
type SlimParticipation struct {
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

// History is a collection of simplified participations
type History []SlimParticipation

// SlimParticipationByYear is a slim participation indexed by participation year to simplify resarch
type SlimParticipationByYear map[int]SlimParticipation

// HistoryByYear is an history indexed by participation year to simplify resarch
type HistoryByYear map[int]History

// HistoryRepository is the interface for the persistency container
type HistoryRepository interface {
	FindByPlayerID(ctx context.Context, id int) (SlimParticipationByYear, error)
	FindByPlayerIDOrderByYear(ctx context.Context, id int) (HistoryByYear, []int, error)
}

// HistorySwitcheroo(ctx context.Context, oldID, newID int, newHistory []History) error

// IsEqual compares two slimParticipations and returns true if they are equivalent,
// false otherwise
func (s *SlimParticipation) IsEqual(b SlimParticipation) bool {
	return s.Year == b.Year && s.Category == b.Category && s.City == b.City && s.IsParis == b.IsParis
}
