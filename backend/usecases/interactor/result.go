package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
)

// CalculateMaxScoreObtainable calculates  the maximum score obtainable in a category for the year
func (i *Interactor) CalculateMaxScoreObtainable(game domain.Game) (int, error) {
	var max int
	start, err := i.GameRepository.FindStartByYearAndCategory(context.Background())
	if err != nil {
		return max, err
	}
	end, err := i.GameRepository.FindEndByYearAndCategory(game.Year, game.Category)
	if err != nil {
		return max, err
	}

	for i := start; i <= end; i++ {
		max += i
	}

	return max, nil
}
