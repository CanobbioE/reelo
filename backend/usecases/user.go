package usecases

import "context"

// User represents the user entity as it is stored in the db
type User struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// UserRepository is the interface for the persistency container
type UserRepository interface {
	FindPasswordByUsername(ctx context.Context, u string) (string, error)
}
