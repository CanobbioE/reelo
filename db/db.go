package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var ctx = context.Background()

func DB() *sql.DB {
	db, err := sql.Open("driver", "dataSourceName")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	return db
}

func ContainsPlayer(db *sql.DB, name, surname string) bool {
	q := `
	SELECT id
	FROM giocatore
	WHERE nome = ? AND cognome = ?
	`
	rows, err := db.QueryContext(ctx, q, name, surname)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

// TODO
func Add(db *sql.DB, table string, params ...interface{}) int {
	var id int
	switch table {
	case "giocatore":
		// TODO
		name := params[0].(string)
		surname := params[1].(string)
		fmt.Println(name, surname)
	case "risultato":
		// TODO
	case "partecipazione":
		// TODO
	case "giochi":
		// TODO
	}
	return id
}
