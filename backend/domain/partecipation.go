package domain

import "context"

// Partecipation represents the partecipation relation as it is stored in the db
type Partecipation struct {
	ID     int    `json:"id"`
	Player Player `json:"player"`
	Game   Game   `json:"game"`
	Result Result `json:"result"`
	City   string `json:"city"`
}

// PartecipationRepository is the interface for the persistency container
type PartecipationRepository interface {
	Store(ctx context.Context, p Partecipation) (int64, error)
	FindCitiesByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) ([]string, error)
	FindAll(ctx context.Context, page, size int) ([]Partecipation, error)
	UpdatePlayerIDByGameID(ctx context.Context, pid, gid int) error
}
