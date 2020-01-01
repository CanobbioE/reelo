package db

import (
	"context"
	"fmt"
	"log"
)

// ContainsPlayer verifies if a player is already in the database
func (database *DB) ContainsPlayer(ctx context.Context, name, surname string) bool {
	q := `SELECT id FROM Giocatore WHERE nome = ? AND cognome = ?`
	rows, err := database.db.QueryContext(ctx, q, name, surname)
	if err != nil {
		log.Printf("Error cchecking player's existence: %v\n", err)
		return false
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

// performAndReturn performs two queries (q1 and q2).
// The first one inserts a row in the DB and the second one gets the ID of the
// inserted row. The id is then returned.
func (database *DB) performAndReturn(ctx context.Context, q1, q2 string) (int64, error) {
	// TODO: everything here is terrible btw
	result, err := database.db.ExecContext(ctx, q1)
	if err != nil {
		return -1, fmt.Errorf("error performing query:\n%s\n%v", q1, err)
	}
	// Try to get the id with the built in function
	id, err := result.LastInsertId()
	if err != nil {
		if q2 != "" {
			// Try to get the id with a query
			err := database.db.QueryRowContext(ctx, q2).Scan(&id)
			if err != nil {
				return -1, fmt.Errorf("error retrieving ID with:\n%s\n%v", q2, err)
			}
		}
	}
	return id, nil
}

func adaptToParis(query string, isParis bool) string {
	if isParis {
		return fmt.Sprintf("%s\nAND P.sede = \"paris\"\n", query)
	}
	return fmt.Sprintf("%s\nAND P.sede <> \"paris\"\n", query)
}

// CreateAccent returns an identifier based on a player's first year and first city of partecipation
func CreateAccent(year, num int, city string) string {
	return fmt.Sprintf("%d %s %d", year, city, num)
}

// ContainsComment verifies if a comment is already assigned to the given playerID
func (database *DB) ContainsComment(ctx context.Context, playerID int) bool {
	q := `SELECT id FROM Commenti WHERE giocatore = ?`
	rows, err := database.db.QueryContext(ctx, q, playerID)
	if err != nil {
		log.Printf("Error cchecking comment's existence: %v\n", err)
		return false
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

// TODO: make a function for all the single result queries
// TODO: make a function for all the multiple results queries
