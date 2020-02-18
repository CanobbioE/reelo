package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
)

// ListRanks returns a list of all the ranks in the database
func (i *Interactor) ListRanks(page, size int) ([]domain.Partecipation, error) {
	return i.PartecipationRepository.FindAll(context.Background(), page, size)
}
