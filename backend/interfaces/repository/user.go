package repository

import (
	"context"
	"fmt"
)

// USERREPO is the handler name
const USERREPO = "UserRepo"

// DbUserRepo id the repository for Users
type DbUserRepo DbRepo

// NewDbUserRepo istanciates and returns a User repository
func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
	return &DbUserRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[UserREPO],
	}
}

// FindPasswordByUsername retrieves a user's password given its username.
// Expects the password to be hashed
func (db *DbUserRepo) FindPasswordByUsername(ctx context.Context, u string) (string, error) {
	var hash string
	q := `SELECT parolachiave FROM Utenti WHERE nomeutente = %s`
	q = fmt.Sprintf(q, u)
	err := QueryRow(ctx, q, db.dbHandler, &hash)
	return hash, err
}
