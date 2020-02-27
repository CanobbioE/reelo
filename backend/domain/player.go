package domain

import "context"

// Player represents the player entity as it is stored in the db
type Player struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Reelo   float64 `json:"reelo"`
	Accent  string  `json:"accent"`
}

// PlayerRepository is the interface for the persistency container
type PlayerRepository interface {
	Store(ctx context.Context, p Player) (int64, error)
	DeleteByID(ctx context.Context, id int) error
	FindIDByNameAndSurname(ctx context.Context, n, s string) (int, error)
	FindAllIDs(ctx context.Context) ([]int, error)
	FindByID(ctx context.Context, id int) (Player, error)
	FindAll(ctx context.Context, page, size int) ([]Player, error)
	FindAllOrderByReeloDesc(ctx context.Context, page, size int) ([]Player, error)
	FindCountAll(ctx context.Context) (int, error)
	UpdateReelo(ctx context.Context, p Player) error
	UpdateAccent(ctx context.Context, p Player) error
	CheckExistenceByNameAndSurname(ctx context.Context, n, s string) bool
	FindAllIDsWhereIDNotInPartecipation(ctx context.Context) ([]int, error)
}
