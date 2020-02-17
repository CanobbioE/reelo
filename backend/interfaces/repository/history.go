package repository

import (
	"context"

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

// FindByPlayerID retrieves the history for the given player id
func (db *DbHistoryRepo) FindByPlayerID(ctx context.Context, id int) (usecases.History, []int, error) {
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

	/*
		type History map[int]struct {
			Partecipation domain.Partecipation `json:"partecipation"`
			MaxExercises  int                  `json:"eMax"`
			MaxScore      int                  `json:"dMax"`
		}
	*/
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
