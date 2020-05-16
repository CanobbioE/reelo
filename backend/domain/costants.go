package domain

import "context"

// Costants represents some of the values used to calculate a player's Reelo.
// These are the values that can be changed to tinker with the algorithm.
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
