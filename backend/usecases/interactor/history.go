package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
)

// PlayerHistory returns the history details for the player
func (i *Interactor) PlayerHistory(player domain.Player) (usecases.SlimPartecipationByYear, error) {
	return i.HistoryRepository.FindByPlayerID(context.Background(), player.ID)
}

// AnalysisHistory returns an analysis history
func (i *Interactor) AnalysisHistory(player domain.Player) (usecases.HistoryByYear, []int, error) {
	return i.HistoryRepository.FindByPlayerIDOrderByYear(context.Background(), player.ID)
}
