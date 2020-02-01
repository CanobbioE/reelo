package domain

import "context"

// Comment represents the comment entity as it is stored in the db
type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Player Player `json:"player"`
}

// CommentRepository is the interface for the persistency container
type CommentRepository interface {
	FindByID(ctx context.Context, id int) (string, error)
	UpdateTextByPlayerID(ctx context.Context, id int, t string)
	CheckExistenceByPlayerID(ctx context.Context, id int) bool
}
