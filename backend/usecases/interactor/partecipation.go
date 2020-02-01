package interactor

import "context"

// ListRanks returns a list of all the ranks in the database
func (i *Interactor) ListRanks(page, size int) ([]Partecipation, error) {
	return i.PartecipationRepository.FindAll(context.Background(), page, size)
}
