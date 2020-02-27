package repository

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/domain"
)

// RESULTREPO is the handler name
const RESULTREPO = "ResultRepo"

// DbResultRepo id the repository for Results
type DbResultRepo DbRepo

// NewDbResultRepo istanciates and returns a Result repository
func NewDbResultRepo(dbHandlers map[string]DbHandler) *DbResultRepo {
	return &DbResultRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[RESULTREPO],
	}
}

// Store a new result entity in the repository
func (db *DbResultRepo) Store(ctx context.Context, r domain.Result) (int64, error) {
	s := `INSERT INTO Risultato (tempo, esercizi, punteggio, posizione, pseudo_reelo)
 			VALUES (%d, %d, %d, %d, %f)`
	s = fmt.Sprintf(s, r.Time, r.Exercises, r.Score, r.Position, r.PseudoReelo)

	result, err := db.dbHandler.ExecContext(ctx, s)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

// FindAllByPlayerID retrives all the results a player had in all the years he partecipated
func (db *DbResultRepo) FindAllByPlayerID(ctx context.Context, id int) ([]domain.Result, error) {
	var results []domain.Result
	// TODO we are not using category and year
	q := `SELECT R.tempo, R.esercizi, R.punteggio, G.anno, G.categoria
			FROM Giocatore U
			JOIN Partecipazione P ON P.giocatore = U.id
			JOIN Risultato R ON R.id = P.risultato
			JOIN Giochi G ON G.id = P.giochi
			WHERE U.id = ?
			`
	rows, err := db.dbHandler.Query(ctx, q, id)
	if err != nil {
		return results, fmt.Errorf("Error getting results: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var r domain.Result
		err := rows.Scan(r.Time, r.Exercises, r.Score)
		if err != nil {
			return results, fmt.Errorf("Error getting results: %v", err)
		}
		results = append(results, r)
	}
	return results, nil

}

// FindExercisesByPlayerID does stuff
func (db *DbResultRepo) FindExercisesByPlayerID(ctx context.Context, id int) (int, error) {
	// TODO

	return 0, nil
}

// FindScoreByYearAndPlayerIDAndGameIsParis retrieve the score of a given player for a given year
func (db *DbResultRepo) FindScoreByYearAndPlayerIDAndGameIsParis(ctx context.Context, y, id int, ip bool) (float64, error) {
	var score float64
	q := `SELECT R.punteggio FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giochi G ON G.id = P.giochi
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = ? AND G.anno = ?`

	q = adaptToParis(q, ip)

	q = fmt.Sprintf(q, id, y)
	err := QueryRow(ctx, q, db.dbHandler, &score)
	return score, err

}

// FindExercisesByYearAndPlayerIDAndGameIsParis does stuff
func (db *DbResultRepo) FindExercisesByYearAndPlayerIDAndGameIsParis(ctx context.Context, y, id int, ip bool) (int, error) {
	var ex int
	q := `SELECT R.esercizi FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giochi G ON G.id = P.giochi
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = ? AND G.anno = ?`

	q = adaptToParis(q, ip)
	q = fmt.Sprintf(q, id, y)
	err := QueryRow(ctx, q, db.dbHandler, &ex)
	return ex, err

}

// FindAvgScoreByGameYear does stuff
func (db *DbResultRepo) FindAvgScoreByGameYear(ctx context.Context, y, k int) (float64, error) {
	var avg float64
	q := `SELECT AVG(X.avg) FROM (
			SELECT AVG(%d * R.esercizi + R.punteggio) AS avg
			FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giochi G ON G.id = P.giochi
			WHERE G.anno = %d
			GROUP BY G.categoria) AS X`

	q = fmt.Sprintf(q, k, y)
	err := QueryRow(ctx, q, db.dbHandler, &avg)
	return avg, err
}

// FindAvgPseudoReeloByGameYearAndCategory returns the pseudo-Reelo's average for the given year and category
func (db *DbResultRepo) FindAvgPseudoReeloByGameYearAndCategory(ctx context.Context, y int, c string) (float64, error) {
	var avg float64
	q := `SELECT IFNULL(AVG(R.pseudo_reelo), -1) FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giochi G ON G.id = P.giochi
			WHERE G.anno = %d AND G.categoria = "%s"`

	q = fmt.Sprintf(q, y, c)
	err := QueryRow(ctx, q, db.dbHandler, &avg)
	return avg, err
}

// FindMaxScoreByGameYearAndCategory calculates the maximum score obtained by any player
// in the given year and category
func (db *DbResultRepo) FindMaxScoreByGameYearAndCategory(ctx context.Context, y int, c string) (float64, error) {
	var max float64
	q := `SELECT IFNULL(MAX(R.punteggio), -1) FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giochi G ON G.id = P.giochi
			WHERE G.anno = %d AND G.categoria = "%s"`

	q = fmt.Sprintf(q, y, c)
	err := QueryRow(ctx, q, db.dbHandler, &max)
	return max, err
}

// FindPseudoReeloByPlayerIDAndGameYear retreives a pseudo-Reelo for a given player in the specified year
func (db *DbResultRepo) FindPseudoReeloByPlayerIDAndGameYear(ctx context.Context, id, y int) (float64, error) {
	var pseudoReelo float64
	q := `SELECT R.pseudo_reelo FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giocatore U ON U.id = P.giocatore
			JOIN Giochi G ON G.id = P.giochi
			WHERE U.id = %d AND G.anno = %d`

	q = fmt.Sprintf(q, id, y)

	err := QueryRow(ctx, q, db.dbHandler, &pseudoReelo)
	return pseudoReelo, err
}

// FindIDByPlayerIDAndGameYearAndCategory eturns the id of a result given a player a year and a category
func (db *DbResultRepo) FindIDByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) (int, error) {
	var resultID int

	q := `SELECT R.id
			FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giocatore U ON U.id = P.Giocatore
			JOIN Giochi G ON G.id = P.giochi
			WHERE U.id = %d AND G.anno = %d AND G.categoria = "%s"`
	q = fmt.Sprintf(q, id, y, c)

	err := QueryRow(ctx, q, db.dbHandler, &resultID)
	return resultID, err
}

// UpdatePseudoReeloByPlayerIDAndGameYearAndCategory sets a new value for the
// pseudo reelo identified by the given player id, game year and game category
func (db *DbResultRepo) UpdatePseudoReeloByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string, pr float64) error {
	resultID, err := NewDbResultRepo(db.dbHandlers).FindIDByPlayerIDAndGameYearAndCategory(ctx, id, y, c)
	if err != nil {
		return err
	}

	s := `UPDATE Risultato SET pseudo_reelo = %d  WHERE id = %d`
	s = fmt.Sprintf(s, pr, resultID)
	_, err = db.dbHandler.ExecContext(ctx, s)
	return err
}

// DeleteByGameID remove a game entity from the repository, given its id
func (db *DbResultRepo) DeleteByGameID(ctx context.Context, id int) error {
	s := `DELETE FROM Giochi WHERE id = %d`
	s = fmt.Sprintf(s, id)
	_, err := db.dbHandler.ExecContext(ctx, s)
	return err
}

// FindByPlayerIDAndGameYear retrieves a result given a player id and a game's year
func (db *DbResultRepo) FindByPlayerIDAndGameYear(ctx context.Context, id, y int) (domain.Result, error) {
	// G.categoria
	var r domain.Result
	q := `SELECT R.tempo, R.esercizi, R.punteggio, R.pseudo_reelo, R.posizione
			FROM  Giochi G
			JOIN Partecipazione P ON P.giochi = G.id
			JOIN Risultato R ON R.id = P.risultato
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = %d AND G.anno = %d`
	q = fmt.Sprintf(q, id, y)

	err := QueryRow(ctx, q, db.dbHandler, &r.Time, &r.Exercises, &r.Score, &r.PseudoReelo, &r.Position)

	return r, err
}
