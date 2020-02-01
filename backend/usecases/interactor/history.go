package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
)

// PlayerHistory returns the history details for the player
func (i *Interactor) PlayerHistory(player domain.Player) (usecases.History, []int, error) {
	return i.HistoryRepository.FindByPlayerID(context.Background, player.ID)

}
