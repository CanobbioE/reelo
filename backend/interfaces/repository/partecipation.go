package repository

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/domain"
)

// PARTECIPATIONREPO is the handler name
const PARTECIPATIONREPO = "PartecipationRepo"

// DbPartecipationRepo id the repository for Partecipations
type DbPartecipationRepo DbRepo

// NewDbPartecipationRepo istanciates and returns a Partecipation repository
func NewDbPartecipationRepo(dbHandlers map[string]DbHandler) *DbPartecipationRepo {
	return &DbPartecipationRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[PARTECIPATIONREPO],
	}
}

// Store creates a new partecipation relationship in the repository
func (db *DbPartecipationRepo) Store(ctx context.Context, p domain.Partecipation) (int64, error) {
	s := `INSERT INTO Partecipazione (giocatore, giochi, risultato, sede)
			VALUES (%d, %d, %d, "%s")`
	s = fmt.Sprintf(s, p.Player.ID, p.Game.ID, p.Result.ID, p.City)
	result, err := db.dbHandler.ExecContext(ctx, s)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// FindCitiesByPlayerIDAndGameYearAndCategory returns a list of cities for
// the given player's ID, game's year and game's category
func (db *DbPartecipationRepo) FindCitiesByPlayerIDAndGameYearAndCategory(ctx context.Context, id, y int, c string) ([]string, error) {
	var cities []string

	q := `SELECT P.sede FROM Partecipazione P
			JOIN Giocatore U ON U.id = P.giocatore
			JOIN Giochi G ON G.id = P.giochi
			WHERE U.id = ?
			AND G.anno = ? AND G.categoria = ?`

	rows, err := db.dbHandler.Query(ctx, q, id, y, c)
	if err != nil {
		return cities, err
	}
	defer rows.Close()

	for rows.Next() {
		var c string
		err := rows.Scan(&c)
		if err != nil {
			return cities, err
		}
		cities = append(cities, c)
	}

	return cities, nil
}

// UpdatePlayerIDByGameID updates all the parteciaptions that contains
// the specified gameID by changing the player ID to the specified one
func (db *DbPartecipationRepo) UpdatePlayerIDByGameID(ctx context.Context, pid, gid int) error {
	q := `UPDATE Partecipazione SET Giocatore = %d  WHERE Giochi = %d`

	q = fmt.Sprintf(q, pid, gid)

	_, err := db.dbHandler.ExecContext(ctx, q)
	return err
}
