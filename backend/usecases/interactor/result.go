package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
)

// CalculateMaxScoreObtainable calculates  the maximum score obtainable in a category for the year
func (i *Interactor) CalculateMaxScoreObtainable(game domain.Game) (int, utils.Error) {
	var max int
	start, err := i.GameRepository.FindStartByYearAndCategory(context.Background(), game.Year, game.Category)
	if err != nil {
		i.Logger.Log("CalculateMaxScoreObtainable: cannot find starting exercise: %v", err)
		return max, utils.NewError(err, "E_GENERIC", 500)
	}
	end, err := i.GameRepository.FindEndByYearAndCategory(context.Background(), game.Year, game.Category)
	if err != nil {
		i.Logger.Log("CalculateMaxScoreObtainable: cannot find ending exercise: %v", err)
		return max, utils.NewError(err, "E_GENERIC", 500)
	}

	for i := start; i <= end; i++ {
		max += i
	}

	return max, utils.NewNilError()
}
