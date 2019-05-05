package db

import (
	"context"
	"crypto/sha256"
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

// DB is a wrapper for the sql.DB
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
		log.Fatalf("Error opening the database: %s", err)
	}
	database := DB{
		db: db,
	}
	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	return &database
}

// ContainsPlayer verifies if a player is already in the database
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
	// TODO: everything here is terrible btw
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

// GetResults retrives all the results a player had in all the years he partecipated
func (database *DB) GetResults(ctx context.Context, name, surname string) (results []Result) {
	q := `
	SELECT R.tempo, R.esercizi, R.punteggio, G.anno, G.categoria
	FROM Giocatore U
	JOIN Partecipazione P ON P.giocatore = U.id
	JOIN Risultato R ON R.id = P.risultato
	JOIN Giochi G ON G.id = P.giochi
	WHERE U.Nome = ? AND U.Cognome = ?
`
	rows, err := database.db.QueryContext(ctx, q, name, surname)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Result{}
		err := rows.Scan(&r.Time, &r.Exercises, &r.Score, &r.Year, &r.Category)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, r)
	}
	return results
}

// TODO Since all the following functions recovers information that could be
// obtained by manipulating the Result array returned GetResults,
// we should refactor all of this as a datatype. Be careul to not overdo it:
// using SQL is really efficient compared to iterating over data structures
// avg, min and max functions should be DB based.
type Result struct {
	Time      int
	Exercises int
	Score     int
	Year      int
	Category  string
}

// TODO implement GetScore
func (database *DB) GetScore(name, surname string, year int) int {
	return 0
}

//TODO implement GetExercises
func (database *DB) GetExercises(name, surname string, year int) int {
	return 0
}

// TODO: implement GetCategory
func (database *DB) GetCategory(name, surname string, year int) string {
	return "nope"
}

// TODO: implement GetAvgScoresOfCategories
func (database *DB) GetAvgScoresOfCategories(year int) float64 {
	return 0
}

// TODO: implement GetAvgScore
func (database *DB) GetAvgScore(year int, category string) float64 {
	return 0
}

// TODO: implement GetMaxScore
func (database *DB) GetMaxScore(year int, category string) float64 {
	return 0
}

// TODO: implement GetCreds
func GetCreds(username string) (string, error) {
	if username == "admin-canna" {
		hashPassword := sha256.New()
		hashPassword.Write([]byte("citrosodina"))
		toRet := hashPassword.Sum(nil)
		return string(toRet), nil
	}
	return "", fmt.Errorf("user not found")
}
