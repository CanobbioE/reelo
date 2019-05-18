package db

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/utils/parse"
)

// Add is used to insert a new row in the specified db and table.
// The id of the newly added row is returned for further reference.
// Possible tables are: 'giocatore', 'risultato', 'partecipazione', 'giochi'
func (database *DB) Add(ctx context.Context, table string, params ...interface{}) int {
	var q1, q2 string

	switch table {
	case "giocatore":
		q1 = fmt.Sprintf("INSERT INTO Giocatore (nome, cognome, reelo)"+
			" VALUES (\"%s\", \"%s\", 0)", params...)
		q2 = fmt.Sprintf("SELECT id FROM Giocatore"+
			" WHERE nome = \"%s\" AND cognome = \"%s\"", params...)

	case "risultato":
		q1 = fmt.Sprintf("INSERT INTO Risultato (tempo, esercizi, punteggio)"+
			" VALUES (%d, %d, %d)", params...)

		q2 = fmt.Sprintf("SELECT MAX(id) FROM Risultato "+
			"WHERE tempo = %d AND esercizi = %d AND punteggio = %d", params...)

	case "partecipazione":
		q1 = fmt.Sprintf("INSERT INTO Partecipazione (giocatore, giochi, risultato, sede) "+
			"VALUES (%d, %d, %d, \"%s\")", params...)

	case "giochi":
		q1 = fmt.Sprintf("INSERT INTO Giochi (anno, categoria)"+
			" VALUES (%d, \"%s\")", params...)

		q2 = fmt.Sprintf("SELECT id FROM Giochi"+
			" WHERE anno = %d AND categoria = %s", params...)
	}
	return int(database.performAndReturn(ctx, q1, q2))
}

// InserRankingFile inserts all the result contained in the already parsed file into the database by making the correct calls
func (database DB) InserRankingFile(ctx context.Context,
	file []parse.LineInfo, year int, category string, isParis bool) {
	gamesID := database.Add(ctx, "giochi", year, category)
	for _, line := range file {
		city := line.City
		if isParis {
			city = "paris"
		}
		var playerID int
		if !database.ContainsPlayer(ctx, line.Name, line.Surname) {
			playerID = database.Add(ctx, "giocatore", line.Name, line.Surname)
		}
		playerID = database.PlayerID(ctx, line.Name, line.Surname)
		resultsID := database.Add(ctx, "risultato", line.Time, line.Exercises, line.Points)
		database.Add(ctx, "partecipazione", playerID, gamesID, resultsID, city)
	}
}

// UpdateReelo sets a new reelo for the specified player
func (database *DB) UpdateReelo(ctx context.Context, p Player) error {
	q := `UPDATE Giocatore SET reelo = ? WHERE nome = ? AND cognome = ?`
	_, err := database.db.ExecContext(ctx, q, p.Reelo, p.Name, p.Surname)
	if err != nil {
		return err
	}
	return nil
}
