package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// PlayerHistory returns the history details for the player.
// TODO: soon obsolete
func (i *Interactor) PlayerHistory(player domain.Player) (usecases.SlimParticipationByYear, utils.Error) {
	history, err := i.HistoryRepository.FindByPlayerID(context.Background(), player.ID)
	if err != nil {
		i.Log("PlayerHistory: cannot find for player %v: %v", player.ID, err)
		return history, utils.NewError(err, "E_DB_FIND", 500)
	}
	return history, utils.NewNilError()
}

// AnalysisHistory returns an analysis history.
// TODO: soon obsolete
func (i *Interactor) AnalysisHistory(player domain.Player) (usecases.HistoryByYear, []int, utils.Error) {
	history, years, err := i.HistoryRepository.FindByPlayerIDOrderByYear(context.Background(), player.ID)
	if err != nil {
		i.Log("AnalysisHistory: cannot find for player %v: %v", player.ID, err)
		return history, years, utils.NewError(err, "E_DB_FIND", 500)
	}
	return history, years, utils.NewNilError()
}
