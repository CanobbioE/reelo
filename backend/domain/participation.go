package domain

import "context"

// Participation represents the participation relation as it is stored in the db
type Participation struct {
	Player Player `json:"player"`
	Game   Game   `json:"game"`
	Result Result `json:"result"`
	City   string `json:"city"`
}

// ParticipationRepository is the interface for the persistency container
type ParticipationRepository interface {
	Store(ctx context.Context, p Participation) (int64, error)
	FindByPlayerID(ctx context.Context, id int) ([]Participation, error)
	FindCitiesByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) ([]string, error)
	UpdatePlayerIDByResultID(ctx context.Context, pid, rgid int) error
}
