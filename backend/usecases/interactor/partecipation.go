package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
)

// ListRanks returns a list of all the ranks in the database
func (i *Interactor) ListRanks(page, size int) ([]domain.Partecipation, error) {
	var partecipations []domain.Partecipation
	players, err := i.PlayerRepository.FindAllOrderByReeloDesc(context.Background(), page, size)
	if err != nil {
		return partecipations, err
	}

	for _, player := range players {
		lastCategory, err := i.GameRepository.FindMaxCategoryByPlayerID(context.Background(), player.ID)
		if err != nil {
			return partecipations, err
		}

		partecipations = append(partecipations, domain.Partecipation{
			Game: domain.Game{
				Category: lastCategory,
			},
			Player: player,
		})
	}

	return partecipations, nil
}
