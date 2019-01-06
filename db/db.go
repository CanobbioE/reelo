package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER = "mysql"
	DB_USER   = "reeloUser"
	DB_PASS   = "password"
	DB_NAME   = "reelo"
	DB_HOST   = "localhost"
)

var dataSourceName = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME

type DB struct {
	db *sql.DB
}

// Close closes the database
func (databse *DB) Close() {
	databse.db.Close()
}

// NewDB returns the databse used for this program.
// REMEMBER TO CLOSE IT!
func NewDB() *DB {
	db, err := sql.Open(DB_DRIVER, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	database := DB{
		db: db,
	}
	if err := db.PingContext(context.Background()); err != nil {
		log.Fatal(err)
	}
	return &database
}

// ContainsPlayer verufies if a player is already in the database
func (database *DB) ContainsPlayer(ctx context.Context, name, surname string) bool {
	q := `
	SELECT id
	FROM giocatore
	WHERE nome = ? AND cognome = ?
	`
	rows, err := database.db.QueryContext(ctx, q, name, surname)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

// Add is used to insert a new row in the specified db and table.
// The id of the newly added row is returned for further reference.
// Possible tables are: 'giocatore', 'risultato', 'partecipazione', 'giochi'
func (database *DB) Add(ctx context.Context, table string, params ...interface{}) int {
	var q1, q2, s string

	switch table {
	case "giocatore":
		s = `
INSERT INTO Giocatore (nome, cognome, reelo)
VALUES (%s, %s, 0)
`
		q1 = fmt.Sprintf(s, params)
		s = `
SELECT id FROM Giocatore
WHERE nome = %s AND cognome = %s
`
		q2 = fmt.Sprintf(s, params)

	case "risultato":
		s = `
INSERT INTO Risultato (tempo, esercizi, punteggio)
VALUES (%d, %d, %d)
`
		q1 = fmt.Sprintf(s, params)

		s = `
SELECT MAX(id) FROM Risultato
WHERE tempo = %d AND esercizi = %d AND punteggio = %d
`
		q2 = fmt.Sprintf(s, params)

	case "partecipazione":
		s = `
INSERT INTO Partecipazione (giocatore, giochi, risultato, sede)
VALUES (%d, %d, %d, %s)
		`
		q1 = fmt.Sprintf(s, params)

	case "giochi":
		s = `
INSERT INTO Giochi (anno, categoria)
VALUES (%d, %s)
		`
		q1 = fmt.Sprintf(s, params)

		s = `
SELECT id FROM Giochi
WHERE anno = %d AND categoria = %s
		`
		q2 = fmt.Sprintf(s, params)
	}
	return int(database.performAndReturn(ctx, q1, q2))
}

// performAndReturn performs two queries (q1 and q2).
// The first one inserts a row in the DB and the second one gets the ID of the
// inserted row. The id is then returned.
func (database *DB) performAndReturn(ctx context.Context, q1, q2 string) int64 {
	// everything here is terrible btw
	result, err := database.db.ExecContext(ctx, q1)
	if err != nil {
		log.Fatal(err)
	}
	// Try to get the id with the built in function
	id, err := result.LastInsertId()
	if err != nil {
		if q2 != "" {
			// Try to get the id with a query
			err := database.db.QueryRowContext(ctx, q2).Scan(&id)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return id
}

func (database *DB) RetrievePlayerID(ctx context.Context, name, surname string) int {
	var pID int
	q := `
		SELECT id FROM Giocatore
		WHERE nome = ? AND cognome = ?
		`
	err := database.db.QueryRow(q, name, surname).Scan(&pID)
	if err != nil {
		log.Fatal(err)
	}
	return pID
}
