package domain

import "context"

// Costants represent all the possible variables used when calculating the reelo
type Costants struct {
	ID                     int     `json:"id"`
	StartingYear           int     `json:"year,omitempty"`
	ExercisesCostant       float64 `json:"ex,omitempty"`
	PFinal                 float64 `json:"final,omitempty"`
	MultiplicativeFactor   float64 `json:"mult,omitempty"`
	AntiExploit            float64 `json:"exp,omitempty"`
	NoPartecipationPenalty float64 `json:"np,omitempty"`
}

// CostantsRepository is the interface for the persistency container
type CostantsRepository interface {
	FindAll(ctx context.Context) (Costants, error)
	UpdateAll(ctx context.Context, c Costants) error
}
