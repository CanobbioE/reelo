package domain

import "context"

// Comment represents the comment entity as it is stored in the db.
// A comment is used to add additional details to a player.
type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Player Player `json:"player"`
}

// CommentRepository is the interface for the persistency container
type CommentRepository interface {
	Store(ctx context.Context, c Comment) (int64, error)
	FindByPlayerID(ctx context.Context, id int) (Comment, error)
	FindTextByPlayerID(ctx context.Context, id int) (string, error)
	UpdateTextByPlayerID(ctx context.Context, t string, id int) error
	CheckExistenceByPlayerID(ctx context.Context, id int) bool
}
