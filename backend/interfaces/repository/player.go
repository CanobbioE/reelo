package repository

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/domain"
)

// PLAYERREPO is the handler name
const PLAYERREPO = "playerRepo"

// DbPlayerRepo id the repository for Players
type DbPlayerRepo DbRepo

// NewDbPlayerRepo istanciates and returns a Player repository
func NewDbPlayerRepo(dbHandlers map[string]DbHandler) *DbPlayerRepo {
	return &DbPlayerRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[PLAYERREPO],
	}
}

// Store creates a new player entity in the repository
func (db *DbPlayerRepo) Store(ctx context.Context, p domain.Player) (int64, error) {
	s := `INSERT INTO Giocatore (nome, cognome, accent, reelo)
			VALUES ("%s", "%s", "%s", 0)`
	s = fmt.Sprintf(s, p.Name, p.Surname, p.Accent, p.Reelo)

	result, err := db.dbHandler.Execute(ctx, s)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// FindIDByNameAndSurname retrieves a player id from the database given its name and surname
// If an error occurs or the player does not exists
func (db *DbPlayerRepo) FindIDByNameAndSurname(ctx context.Context, n, s string) (int, error) {
	id := -1
	q := `SELECT id FROM Giocatore WHERE nome = %s AND cognome = %s`

	q = fmt.Sprintf(q, n, s)
	err := QueryRow(ctx, q, db.dbHandler, &id)
	return id, err
}

// FindAllIDs retrieves all the players' IDs
func (db *DbPlayerRepo) FindAllIDs(ctx context.Context) ([]int, error) {
	var ids []int
	q := `SELECT id FROM Giocatore`

	rows, err := db.dbHandler.Query(ctx, q)
	if err != nil {
		return ids, fmt.Errorf("Error getting all players id: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return ids, fmt.Errorf("Error getting all players id: %v", err)
		}
		// TODO: this is gonna break once the array becomes too big... YOU NEED PAGINATION!
		ids = append(ids, id)
	}
	return ids, nil
}

// FindByID retrieve a Player given its ID
func (db *DbPlayerRepo) FindByID(ctx context.Context, id int) (domain.Player, error) {
	player := domain.Player{ID: id}
	q := `SELECT nome, cognome, accent, reelo FROM Giocatore WHERE id = %d`
	q = fmt.Sprintf(q, id)

	err := QueryRow(ctx, q, db.dbHandler, &player.Name, &player.Surname, &player.Accent, &player.Reelo)
	return player, err
}

// FindAll retrieves all players from the specified page of the set size
// if the size is a negative number then all the players will be retrieved
func (db *DbPlayerRepo) FindAll(ctx context.Context, page, size int) ([]domain.Player, error) {
	var players []domain.Player
	q := `SELECT nome, cognome, reelo, accent, id FROM Giocatore LIMIT ?, ?`
	count, err := db.FindCountAll(ctx)
	if err != nil {
		return players, err
	}
	if size < 0 {
		size = count
	}
	rows, err := db.dbHandler.Query(ctx, q, (page-1)*size, size)
	if err != nil {
		return players, fmt.Errorf("Error getting players: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.Player
		err := rows.Scan(&p.Name, &p.Surname, &p.Reelo, &p.Accent, &p.ID)
		if err != nil {
			return players, fmt.Errorf("Error getting players: %v", err)
		}
		players = append(players, p)
	}
	return players, nil
}

// FindCountAll returns the nomber of tuples in the player's repository
func (db *DbPlayerRepo) FindCountAll(ctx context.Context) (int, error) {
	var count int

	q := `SELECT COUNT(U.id) FROM Giocatore U`
	err := QueryRow(ctx, q, db.dbHandler, &count)
	return count, err
}

// UpdateReelo sets a new reelo for the specified player
func (db *DbPlayerRepo) UpdateReelo(ctx context.Context, p domain.Player) error {
	s := `UPDATE Giocatore SET reelo = %d WHERE id = %d`
	s = fmt.Sprintf(s, p.Reelo, p.ID)
	_, err := db.dbHandler.Execute(ctx, s)
	return err
}

// UpdateAccent sets a new accent for the specified player
func (db *DbPlayerRepo) UpdateAccent(ctx context.Context, p domain.Player) error {
	s := `UPDATE Giocatore SET accent = %s  WHERE id = %d`
	s = fmt.Sprintf(s, p.Accent, p.ID)
	_, err := db.dbHandler.Execute(ctx, s)
	return err
}

// CheckExistenceByNameAndSurname returns true if a player
// with the given name and surname exists in the repository
func (db *DbPlayerRepo) CheckExistenceByNameAndSurname(ctx context.Context, n, s string) bool {
	q := `SELECT id FROM Giocatore WHERE nome = ? AND cognome = ?`
	rows, err := db.dbHandler.Query(ctx, q, n, s)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		return true
	}
	return false
}
