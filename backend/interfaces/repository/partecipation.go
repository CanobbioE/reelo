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

// UpdatePlayerIDByResultID updates all the parteciaptions that contains
// the specified result's ID by changing the player ID to the specified one
func (db *DbPartecipationRepo) UpdatePlayerIDByResultID(ctx context.Context, pid, rid int) error {
	q := `UPDATE Partecipazione SET giocatore = %d  WHERE risultato = %d`

	q = fmt.Sprintf(q, pid, rid)

	_, err := db.dbHandler.ExecContext(ctx, q)
	return err
}

// FindByPlayerID retrieve all the partecipations that include the given
// player's ID. The sub-structs are populated only with the IDs
func (db *DbPartecipationRepo) FindByPlayerID(ctx context.Context, id int) ([]domain.Partecipation, error) {
	var partecipations []domain.Partecipation
	q := `SELECT giocatore, giochi, risultato, sede
			FROM Partecipazione
			WHERE giocatore = ?`
	rows, err := db.dbHandler.Query(ctx, q, id)
	if err != nil {
		return partecipations, err
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.Partecipation
		err := rows.Scan(&p.Player.ID, &p.Game.ID, &p.Result.ID, &p.City)
		if err != nil {
			return partecipations, err
		}

		partecipations = append(partecipations, p)
	}
	return partecipations, nil
}
