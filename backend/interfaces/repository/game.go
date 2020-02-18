package repository

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/domain"
)

// GAMEREPO is the handler name
const GAMEREPO = "gameRepo"

// DbGameRepo id the repository for Games
type DbGameRepo DbRepo

// NewDbGameRepo istanciates and returns a Game repository
func NewDbGameRepo(dbHandlers map[string]DbHandler) *DbGameRepo {
	return &DbGameRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[GAMEREPO],
	}
}

// Store creates a new game entity in the repository
func (db *DbGameRepo) Store(ctx context.Context, g domain.Game) (int64, error) {
	s := `INSERT INTO Giochi (anno, categoria, inizio, fine, internazionale)
 VALUES (%d, "%s", %d, %d, %t)`

	s = fmt.Sprintf(s, g.Year, g.Category, g.Start, g.End, g.IsParis)
	result, err := db.dbHandler.ExecContext(ctx, s)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// FindIDByYearAndCategoryAndIsParis retrieves the id of the game from the specified year and category.
// If no result are found the id will be -1.
func (db *DbGameRepo) FindIDByYearAndCategoryAndIsParis(ctx context.Context, y int, c string, ip bool) (int, error) {
	id := -1
	q := `SELECT G.id FROM Giochi G
WHERE G.anno = ? AND G.categoria = ? AND G.internazionale = ?`

	q = fmt.Sprintf(q, y, c, ip)
	err := QueryRow(ctx, q, db.dbHandler, &id)
	return id, err
}

// FindDistinctYearsByPlayerID retrieves a list of all the years a player has played
func (db *DbGameRepo) FindDistinctYearsByPlayerID(ctx context.Context, id int) ([]int, error) {
	var years []int

	pID, err := NewDbPlayerRepo(db.dbHandlers).FindByID(ctx, id)
	if err != nil {
		return years, err
	}

	q := `SELECT DISTINCT G.anno FROM Giochi G
JOIN Partecipazione P ON P.giochi = G.id
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = ?`

	rows, err := db.dbHandler.Query(ctx, q, pID)
	if err != nil {
		return years, fmt.Errorf("Error getting partcipation years: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var y int
		err := rows.Scan(&y)
		if err != nil {
			return years, fmt.Errorf("Error getting partcipation years: %v", err)
		}
		years = append(years, y)
	}
	return years, nil
}

// FindCategoriesByYearAndPlayer does stuff
func (db *DbGameRepo) FindCategoriesByYearAndPlayer(ctx context.Context, y, id int) ([]string, error) {
	var categories []string
	q := `SELECT R.punteggio FROM Risultato R
			JOIN Partecipazione P ON P.risultato = R.id
			JOIN Giochi G ON G.id = P.giochi
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = ? AND G.anno = ?`

	rows, err := db.dbHandler.Query(ctx, q, id, y)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var c string
		err := rows.Scan(&c)
		if err != nil {
			return categories, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// FindMaxYear  returns the last year stored in the database
func (db *DbGameRepo) FindMaxYear(ctx context.Context) (int, error) {
	var max int
	q := `SELECT MAX(anno) FROM Giochi`

	err := QueryRow(ctx, q, db.dbHandler, &max)
	return max, err
}

// FindMaxCategoryByPlayerID returns the last category a player has played into
func (db *DbGameRepo) FindMaxCategoryByPlayerID(ctx context.Context, id int) (string, error) {
	var max string
	q := `SELECT G.categoria FROM Giochi G
			JOIN Partecipazione P ON P.giochi = G.id
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = %d
			AND G.anno = (
				SELECT MAX(G.anno) FROM Giochi G
				JOIN Partecipazione P ON P.giochi = G.id
				JOIN Giocatore U ON U.id = P.giocatore
				WHERE U.id = %d
			)`

	q = fmt.Sprintf(q, id, id)
	err := QueryRow(ctx, q, db.dbHandler, &max)
	return max, err
}

// FindStartByYearAndCategory returns the number of the first (a.k.a. the starting)
// exercise for the specified year and category
func (db *DbGameRepo) FindStartByYearAndCategory(ctx context.Context, y int, c string) (int, error) {
	var start int
	q := `SELECT inizio FROM Giochi WHERE anno = %d  AND categoria = %s`
	q = fmt.Sprintf(q, y, c)

	err := QueryRow(ctx, q, db.dbHandler, &start)
	return start, err
}

// FindEndByYearAndCategory returns the number of the last (a.k.a. the ending)
// exercise for the specified year and category
func (db *DbGameRepo) FindEndByYearAndCategory(ctx context.Context, y int, c string) (int, error) {
	var end int
	q := `SELECT fine FROM Giochi WHERE anno = %d  AND categoria = %s`
	q = fmt.Sprintf(q, y, c)

	err := QueryRow(ctx, q, db.dbHandler, &end)
	return end, err
}

// FindCategoryByPlayerIDAndGameYear retreives the category the given player
// has partecipated into, during the specified year
func (db *DbGameRepo) FindCategoryByPlayerIDAndGameYear(ctx context.Context, id, y int) (string, error) {
	var category string
	q := `SELECT G.categoria FROM Giochi G
			JOIN Partecipazione P ON P.giochi = G.id
			JOIN Giocatore U ON U.id = P.giocatore
			WHERE U.id = %d AND G.anno = %d`

	q = fmt.Sprintf(q, id, y)

	err := QueryRow(ctx, q, db.dbHandler, &category)
	return category, err
}

// FindAllYears return a list of all the stored years
func (db *DbGameRepo) FindAllYears(ctx context.Context) ([]int, error) {
	var years []int

	q := `SELECT DISTINCT anno FROM Giochi`

	rows, err := db.dbHandler.Query(ctx, q)
	if err != nil {
		return years, err
	}
	defer rows.Close()

	for rows.Next() {
		var y int
		err := rows.Scan(&y)
		if err != nil {
			return years, err
		}
		years = append(years, y)
	}
	return years, nil
}
