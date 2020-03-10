package repository

import (
	"context"
	"fmt"
	"sort"

	"github.com/CanobbioE/reelo/backend/usecases"
)

// HISTORYREPO is the handler name
const HISTORYREPO = "historyRepo"

// DbHistoryRepo id the repository for Histories
type DbHistoryRepo DbRepo

// NewDbHistoryRepo instantiates and returns a History repository
func NewDbHistoryRepo(dbHandlers map[string]DbHandler) *DbHistoryRepo {
	return &DbHistoryRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[HISTORYREPO],
	}
}

// FindByPlayerIDOrderByYear retrieves the history for the given player id
func (db *DbHistoryRepo) FindByPlayerIDOrderByYear(ctx context.Context, id int) (usecases.HistoryByYear, []int, error) {
	historyByYear := make(usecases.HistoryByYear)
	var years []int
	q := `SELECT G.anno, G.categoria, G.internazionale, P.sede
			FROM Giochi G
			JOIN Partecipazione P ON P.giochi = G.id
			WHERE P.Giocatore = ?
			ORDER BY G.anno`

	rows, err := db.dbHandler.Query(ctx, q, id)
	if err != nil {
		return historyByYear, years, err
	}
	defer rows.Close()

	for rows.Next() {
		var sp usecases.SlimParticipation
		err := rows.Scan(&sp.Year, &sp.Category, &sp.IsParis, &sp.City)
		if err != nil {
			return historyByYear, years, err
		}
		var history []usecases.SlimParticipation
		history, ok := historyByYear[sp.Year]
		if !ok {
			history = make([]usecases.SlimParticipation, 0)
		}
		history = append(history, sp)
		historyByYear[sp.Year] = history
		years = append(years, sp.Year)
	}
	sort.Ints(years)
	return historyByYear, years, nil
}

// FindByPlayerID retrieves the history for the given player id
func (db *DbHistoryRepo) FindByPlayerID(ctx context.Context, id int) (usecases.SlimParticipationByYear, error) {
	ph := make(usecases.SlimParticipationByYear)
	q := `SELECT G.categoria, R.tempo, R.esercizi, R.punteggio, R.pseudo_reelo, R.posizione
			FROM  Giochi G
			JOIN Partecipazione P ON P.giochi = G.id
			JOIN Risultato R ON R.id = P.risultato
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = %d AND G.anno = %d`

	// Find all players participation years
	years, err := NewDbGameRepo(db.dbHandlers).FindDistinctYearsByPlayerID(ctx, id)
	if err != nil {
		return ph, err
	}

	for _, y := range years {
		s := fmt.Sprintf(q, id, y)
		rows, err := db.dbHandler.Query(ctx, s)
		if err != nil {
			return ph, err
		}
		defer rows.Close()

		for rows.Next() {
			var slim usecases.SlimParticipation

			// Saving most of the results
			err := rows.Scan(&slim.Category, &slim.Time,
				&slim.Exercises, &slim.Score,
				&slim.PseudoReelo, &slim.Position)
			if err != nil {
				return ph, err

			}

			dMax, err := NewDbResultRepo(db.dbHandlers).FindMaxScoreByGameYearAndCategory(ctx, y, slim.Category)
			if err != nil {
				return ph, err
			}

			t, err := NewDbGameRepo(db.dbHandlers).FindStartByYearAndCategory(ctx, y, slim.Category)
			if err != nil {
				return ph, err
			}

			n, err := NewDbGameRepo(db.dbHandlers).FindEndByYearAndCategory(ctx, y, slim.Category)
			if err != nil {

				return ph, err
			}

			slim.MaxScore = int(dMax)
			slim.MaxExercises = n - t + 1
			ph[y] = slim
		}
	}

	return ph, nil
}
