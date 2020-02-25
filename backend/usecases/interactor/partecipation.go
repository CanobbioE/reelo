package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
)

// ListRanks returns a list of all the ranks in the database
func (i *Interactor) ListRanks(page, size int) ([]domain.Partecipation, utils.Error) {
	var partecipations []domain.Partecipation
	players, err := i.PlayerRepository.FindAllOrderByReeloDesc(context.Background(), page, size)
	if err != nil {
		i.Logger.Log("ListRanks: cannot find all players: %v", err)
		return partecipations, utils.NewError(err, "E_DB_FIND", 500)
	}

	for _, player := range players {
		lastCategory, err := i.GameRepository.FindMaxCategoryByPlayerID(context.Background(), player.ID)
		if err != nil {
			i.Logger.Log("ListRanks: cannot find max category for player %v: %v", player.ID, err)
			return partecipations, utils.NewError(err, "E_DB_FIND", 500)
		}

		partecipations = append(partecipations, domain.Partecipation{
			Game: domain.Game{
				Category: lastCategory,
			},
			Player: player,
		})
	}

	return partecipations, utils.NewNilError()
}
