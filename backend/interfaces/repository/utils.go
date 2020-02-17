package repository

import (
	"context"
	"fmt"
)

// QueryRow queries for a single row
func QueryRow(ctx context.Context, query string, db DbHandler, dest ...interface{}) error {

	row, err := db.Query(ctx, query)
	if err != nil {
		return err
	}
	defer row.Close()

	err = row.Scan(dest...)
	return err
}

func adaptToParis(query string, isParis bool) string {
	if isParis {
		return fmt.Sprintf("%s\nAND P.sede = \"paris\"\n", query)
	}
	return fmt.Sprintf("%s\nAND P.sede <> \"paris\"\n", query)
}

// All returns all the repositories' name
func All() []string {
	return []string{COMMENTREPO, COSTANTSREPO, GAMEREPO, HISTORYREPO, PARTECIPATIONREPO, PLAYERREPO, RESULTREPO, USERREPO}
}
