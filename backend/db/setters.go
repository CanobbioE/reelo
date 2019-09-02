package db

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/utils/parse"
)

// Add is used to insert a new row in the specified db and table.
// The id of the newly added row is returned for further reference.
// Possible tables are: 'giocatore', 'risultato', 'partecipazione', 'giochi'
func (database *DB) Add(ctx context.Context, table string, params ...interface{}) (int, error) {
	var q1, q2 string

	switch table {
	case "giocatore":
		q1 = fmt.Sprintf("INSERT INTO Giocatore (nome, cognome, reelo)"+
			" VALUES (\"%s\", \"%s\", 0)", params...)
		q2 = fmt.Sprintf("SELECT id FROM Giocatore"+
			" WHERE nome = \"%s\" AND cognome = \"%s\"", params...)

	case "risultato":
		q1 = fmt.Sprintf("INSERT INTO Risultato (tempo, esercizi, punteggio, posizione, pseudo_reelo)"+
			" VALUES (%d, %d, %d, %d, %d)", params...)

		q2 = fmt.Sprintf("SELECT MAX(id) FROM Risultato "+
			"WHERE tempo = %d AND esercizi = %d AND punteggio = %d", params...)

	case "partecipazione":
		q1 = fmt.Sprintf("INSERT INTO Partecipazione (giocatore, giochi, risultato, sede) "+
			"VALUES (%d, %d, %d, \"%s\")", params...)

	case "giochi":
		q1 = fmt.Sprintf("INSERT INTO Giochi (anno, categoria, inizio, fine, internazionale)"+
			" VALUES (%d, \"%s\", %d, %d, %t)", params...)

		q2 = fmt.Sprintf("SELECT id FROM Giochi"+
			" WHERE anno = %d AND categoria = %s AND inizio = %d AND fine = %d AND internazionale = %t", params...)
	}
	id, err := database.performAndReturn(ctx, q1, q2)
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

// InserRankingFile inserts all the result contained in the already parsed file into the database by making the correct calls
func (database DB) InserRankingFile(
	ctx context.Context,
	file []parse.LineInfo,
	gameInfo GameInfo,
	isParis bool) error {

	gamesID, err := database.Add(ctx, "giochi",
		gameInfo.Year, gameInfo.Category, gameInfo.Start, gameInfo.End, isParis)
	if err != nil {
		return err
	}
	for _, line := range file {
		if line.Name == "" && line.Surname == "" {
			continue
		}
		city := line.City
		if isParis {
			city = "paris"
		}
		var playerID int
		if !database.ContainsPlayer(ctx, line.Name, line.Surname) {
			playerID, err = database.Add(ctx, "giocatore", line.Name, line.Surname)
			if err != nil {
				return err
			}
		}
		playerID, err = database.PlayerID(ctx, line.Name, line.Surname)
		if err != nil {
			return err
		}
		resultsID, err := database.Add(ctx, "risultato", line.Time, line.Exercises, line.Points, line.Position, 0)
		if err != nil {
			return err
		}
		_, err = database.Add(ctx, "partecipazione", playerID, gamesID, resultsID, city)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateReelo sets a new reelo for the specified player
func (database *DB) UpdateReelo(ctx context.Context, p Player) error {
	q := `UPDATE Giocatore SET reelo = ? WHERE nome = ? AND cognome = ?`
	_, err := database.db.ExecContext(ctx, q, p.Reelo, p.Name, p.Surname)
	return err
}

// UpdateCostants updates the costants used by the reelo algorithm
func (database *DB) UpdateCostants(ctx context.Context, c Costants) error {
	q := `
UPDATE Costanti SET
anno_inizio = ?,
k_esercizi = ?,
finale = ?,
fattore_moltiplicativo = ?,
exploit = ?,
no_partecipazione = ?`
	_, err := database.db.ExecContext(ctx, q,
		c.StartingYear,
		c.ExercisesCostant,
		c.PFinal,
		c.MultiplicativeFactor,
		c.AntiExploit,
		c.NoPartecipationPenalty)
	return err
}

// UpdatePseudoReelo updates the pseudo reelo associated with the given id
func (database *DB) UpdatePseudoReelo(ctx context.Context,
	name, surname string, year int, category string, pseudoReelo float64) error {
	id, err := database.ResultID(name, surname, year, category)
	if err != nil {
		return err
	}
	q := `
UPDATE Risultato SET
pseudo_reelo = ?
WHERE id = ?`

	_, err = database.db.ExecContext(ctx, q, pseudoReelo, id)
	return err
}

// DeleteResultsFrom deletes all the results from a given year
func (database *DB) DeleteResultsFrom(ctx context.Context, gamesID int) error {
	q := `DELETE FROM Giochi WHERE id = ?`

	_, err := database.db.ExecContext(ctx, q, gamesID)
	return err
}
