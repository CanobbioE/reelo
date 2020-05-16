package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
)

// ListPartecipations returns a list of all the parteciaptions in the database.
func (i *Interactor) ListPartecipations(page, size int) ([]domain.Participation, utils.Error) {
	var participations []domain.Participation
	players, err := i.PlayerRepository.FindAllOrderByReeloDesc(context.Background(), page, size)
	if err != nil {
		i.Logger.Log("ListPartecipations: cannot find all players: %v", err)
		return participations, utils.NewError(err, "E_DB_FIND", 500)
	}

	for _, player := range players {
		lastCategory, err := i.GameRepository.FindMaxCategoryByPlayerID(context.Background(), player.ID)
		if err != nil {
			i.Logger.Log("ListPartecipations: cannot find max category for player %v: %v", player.ID, err)
			return participations, utils.NewError(err, "E_DB_FIND", 500)
		}

		participations = append(participations, domain.Participation{
			Game: domain.Game{
				Category: lastCategory,
			},
			Player: player,
		})
	}

	return participations, utils.NewNilError()
}
