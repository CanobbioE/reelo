package repository

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/domain"
)

// COMMENTREPO is the handler name
const COMMENTREPO = "commentRepo"

// DbCommentRepo id the repository for comments
type DbCommentRepo DbRepo

// NewDbCommentRepo istanciates and returns a comment repository
func NewDbCommentRepo(dbHandlers map[string]DbHandler) *DbCommentRepo {
	return &DbCommentRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[COMMENTREPO],
	}
}

// Store inserts a new comment in the repository
func (db *DbCommentRepo) Store(ctx context.Context, c domain.Comment) (int64, error) {
	s := `INSERT INTO Commenti (giocatore, testo) VALUES (%d, "%s")`
	s = fmt.Sprintf(s, c.Player.ID, c.Text)

	result, err := db.dbHandler.ExecContext(ctx, s)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// FindTextByPlayerID retrieves a comment from the repository, given the id
func (db *DbCommentRepo) FindTextByPlayerID(ctx context.Context, id int) (string, error) {
	var text string
	q := `SELECT testo FROM Commenti WHERE giocatore = %d`
	q = fmt.Sprintf(q, id)

	err := QueryRow(ctx, q, db.dbHandler, &text)
	return text, err
}

// UpdateTextByPlayerID updates a comment defined by the player's id
func (db *DbCommentRepo) UpdateTextByPlayerID(ctx context.Context, t string, id int) error {
	s := `UPDATE Commenti SET testo = "%s" WHERE giocatore = %d`
	s = fmt.Sprintf(s, t, id)
	_, err := db.dbHandler.ExecContext(ctx, s)
	return err
}

// CheckExistenceByPlayerID returns true if the repo contains a comment with for the given player id
func (db *DbCommentRepo) CheckExistenceByPlayerID(ctx context.Context, id int) bool {
	q := `SELECT * FROM Commenti WHERE giocatore = ?`

	rows, err := db.dbHandler.Query(ctx, q, id)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		return true
	}
	return false
}

// FindByPlayerID retrieve a comment given a player id
func (db *DbCommentRepo) FindByPlayerID(ctx context.Context, id int) (domain.Comment, error) {
	var c domain.Comment
	q := `SELECT id, testo FROM Commenti WHERE giocatore = ?`

	row, err := db.dbHandler.Query(ctx, q, id)
	if err != nil {
		return c, err
	}
	err = row.Scan(&c.ID, &c.Text)
	if err != nil {
		return c, err
	}
	return c, nil
}
