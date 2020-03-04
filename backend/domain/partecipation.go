package domain

import "context"

// Partecipation represents the partecipation relation as it is stored in the db
type Partecipation struct {
	Player Player `json:"player"`
	Game   Game   `json:"game"`
	Result Result `json:"result"`
	City   string `json:"city"`
}

// PartecipationRepository is the interface for the persistency container
type PartecipationRepository interface {
	Store(ctx context.Context, p Partecipation) (int64, error)
	FindByPlayerID(ctx context.Context, id int) ([]Partecipation, error)
	FindCitiesByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) ([]string, error)
	UpdatePlayerIDByResultID(ctx context.Context, pid, rgid int) error
}
