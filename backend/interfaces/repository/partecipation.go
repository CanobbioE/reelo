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
	result, err := db.dbHandler.Execute(ctx, s)
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

// FindAll retrieves all the Partecipations in the repository, paginating the results
func (db *DbPartecipationRepo) FindAll(ctx context.Context, page, size int) ([]domain.Partecipation, error) {
	var partecipations []domain.Partecipation

	q := `SELECT
			U.id, U.nome, U.cognome, U.reelo, U.accent,
			G.id, G.anno, G.categoria, G.inizio, G.fine, G.internazionale,
			R.id, R.esercizi, R.tempo, R.punteggio, R.posizione, R.pseudo_reelo,
			P.sede
			FROM Giocatore U
			JOIN Partecipazione P ON P.giocatore = U.id
			JOIN Risultato R ON R.id = P.risultato
			JOIN Giochi G ON G.id = P.giochi
			WHERE (G.anno, U.id) IN (
				SELECT MAX(G.anno), U.id FROM Giochi G
				JOIN Partecipazione P ON P.giochi = G.id
				JOIN Giocatore U ON U.id = P.giocatore
				GROUP BY U.id
			)
			ORDER BY U.reelo DESC
			LIMIT ?, ?`

	rows, err := db.dbHandler.Query(ctx, q, (page-1)*size, size)
	if err != nil {
		return partecipations, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			u    domain.Player
			g    domain.Game
			r    domain.Result
			city string
		)

		err := rows.Scan(
			&u.ID, &u.Name, &u.Surname, &u.Reelo, &u.Accent,
			&g.ID, &g.Year, &g.Category, &g.Start, &g.End, &g.IsParis,
			&r.ID, &r.Exercises, &r.Time, &r.Score, &r.Position, &r.PseudoReelo,
			&city,
		)

		if err != nil {
			return partecipations, err
		}

		p := domain.Partecipation{
			Player: u,
			Game:   g,
			Result: r,
			City:   city,
		}
		partecipations = append(partecipations, p)
	}
	return partecipations, nil
}
