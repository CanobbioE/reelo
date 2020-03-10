package domain

import "context"

// Costants represent all the possible variables used when calculating the reelo
type Costants struct {
	ID                     int     `json:"id"`
	StartingYear           int     `json:"year"`
	ExercisesCostant       float64 `json:"ex"`
	PFinal                 float64 `json:"final"`
	MultiplicativeFactor   float64 `json:"mult"`
	AntiExploit            float64 `json:"exp"`
	NoParticipationPenalty float64 `json:"np"`
}

// CostantsRepository is the interface for the persistency container
type CostantsRepository interface {
	FindAll(ctx context.Context) (Costants, error)
	UpdateAll(ctx context.Context, c Costants) error
}
