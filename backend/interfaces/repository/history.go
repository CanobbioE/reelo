package repository

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/usecases"
)

// HISTORYREPO is the handler name
const HISTORYREPO = "historyRepo"

// DbHistoryRepo id the repository for Historys
type DbHistoryRepo DbRepo

// NewDbHistoryRepo istanciates and returns a History repository
func NewDbHistoryRepo(dbHandlers map[string]DbHandler) *DbHistoryRepo {
	return &DbHistoryRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[HISTORYREPO],
	}
}

// FindByPlayerIDOrderByYear retrieves the history for the given player id
func (db *DbHistoryRepo) FindByPlayerIDOrderByYear(ctx context.Context, id int) (usecases.HistoryByYear, []int, error) {
	history := make(usecases.History)
	var years []int
	q := `SELECT G.anno, G.categoria, G.internazionale, P.sede
			FROM Giochi G
			JOIN Partecipazione P ON P.giochi = G.id
			WHERE P.Giocatore = ?
			ORDER BY G.anno`

	rows, err := db.dbHandler.Query(ctx, q, id)
	if err != nil {
		return history, years, err
	}
	defer rows.Close()

	for rows.Next() {
		// var p domain.Partecipation
		var y int
		years = append(years, y)
		err := rows.Scan(&y)
		if err != nil {
			return history, years, err
		}
	}
	return history, years, nil
}

// FindByPlayerIDAndYear retrieves the history for the given player id
func (db *DbHistoryRepo) FindByPlayerIDAndYear(ctx context.Context, id, y int) (usecases.SlimPartecipationByYear, error) {

	ph := make(usecases.SlimPartecipationByYear)
	q := `SELECT G.categoria, R.tempo, R.esercizi, R.punteggio, R.pseudo_reelo, R.posizione
			FROM  Giochi G
			JOIN Partecipazione P ON P.giochi = G.id
			JOIN Risultato R ON R.id = P.risultato
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = %d AND G.anno = %d`

	q = fmt.Sprintf(q, id, y)
	// Find all players partecipation years
	years, err := NewDbGameRepo(db.dbHandlers).FindDistinctYearsByPlayerID(ctx, id)
	if err != nil {
		return ph, err
	}

	for _, y := range years {
		rows, err := db.dbHandler.Query(ctx, q)
		if err != nil {
			return ph, err
		}
		defer rows.Close()

		for rows.Next() {
			var res usecases.SlimPartecipation

			// Saving most of the results
			err := rows.Scan(&res.Category, &res.Time,
				&res.Exercises, &res.Score,
				&res.PseudoReelo, &res.Position)
			if err != nil {
				return ph, err
			}

			// Finding other cool stuff
			dMax, err := NewDbResultRepo(db.dbHandlers).FindMaxScoreByGameYearAndCategory(ctx, y, res.Category)
			if err != nil {
				return ph, err
			}

			t, err := NewDbGameRepo(db.dbHandlers).FindStartByYearAndCategory(ctx, y, res.Category)
			if err != nil {
				return ph, err
			}

			n, err := NewDbGameRepo(db.dbHandlers).FindEndByYearAndCategory(ctx, y, res.Category)
			if err != nil {
				return ph, err
			}

			res.MaxScore = int(dMax)
			res.MaxExercises = n - t + 1
			ph[y] = res
		}
	}

	return ph, nil
}
