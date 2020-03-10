package repository

import (
	"context"
	"fmt"

	"github.com/CanobbioE/reelo/backend/domain"
)

// COSTANTSREPO is the handler name
const COSTANTSREPO = "costantsRepo"

// DbCostantsRepo id the repository for Costantss
type DbCostantsRepo DbRepo

// NewDbCostantsRepo istanciates and returns a Costants repository
func NewDbCostantsRepo(dbHandlers map[string]DbHandler) *DbCostantsRepo {
	return &DbCostantsRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[COSTANTSREPO],
	}
}

// UpdateAll updtades all the costants, this sounds weird because there is only
// one tuple in the costants repository
func (db *DbCostantsRepo) UpdateAll(ctx context.Context, c domain.Costants) error {
	s := `UPDATE Costanti SET
			anno_inizio = %v,
			k_esercizi = %v,
			finale = %v,
			fattore_moltiplicativo = %v,
			exploit = %v,
			no_partecipazione = %v`
	s = fmt.Sprintf(s,
		c.StartingYear,
		c.ExercisesCostant,
		c.PFinal,
		c.MultiplicativeFactor,
		c.AntiExploit,
		c.NoParticipationPenalty)

	_, err := db.dbHandler.ExecContext(ctx, s)
	return err
}

// FindAll retrieves all the costants in the repository.
// This is confusing due to having only one entry in the "costants" table
func (db *DbCostantsRepo) FindAll(ctx context.Context) (domain.Costants, error) {
	var c domain.Costants
	q := `SELECT
			anno_inizio,
			k_esercizi,
			finale,
			fattore_moltiplicativo,
			exploit,
			no_partecipazione
			FROM Costanti`
	err := QueryRow(ctx, q, db.dbHandler,
		&c.StartingYear,
		&c.ExercisesCostant,
		&c.PFinal,
		&c.MultiplicativeFactor,
		&c.AntiExploit,
		&c.NoParticipationPenalty,
	)
	return c, err
}
