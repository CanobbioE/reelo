package usecases

import "context"

// User represents the user entity as it is stored in the db.
// A user is someone that has an account to access the application's administration panel.
// No new user can be registered via the application.
type User struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// UserRepository is the interface for the persistency container
type UserRepository interface {
	FindPasswordByUsername(ctx context.Context, u string) (string, error)
}
