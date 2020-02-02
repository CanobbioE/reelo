package repository

import "context"

// UserREPO is the handler name
const UserREPO = "UserRepo"

// DbUserRepo id the repository for Users
type DbUserRepo DbRepo

// NewDbUserRepo istanciates and returns a User repository
func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
	return &DbUserRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[UserREPO],
	}
}

// FindPasswordByUsername retrieves a user's password given its username
func (db *DbUserRepo) FindPasswordByUsername(ctx context.Context, u string) (string, error) {
	return "", nil
}
