package webinterface

import (
	"io"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// Interactor define the usecases behavior, using an interface allows
// to change implementation as needed
type Interactor interface {
	AnalysisHistory(player domain.Player) (usecases.HistoryByYear, []int, utils.Error)
	AddComment(c domain.Comment) utils.Error
	CalculateAllReelo() utils.Error
	DeleteIfAlreadyExists(game domain.Game) utils.Error
	DoesRankExist(year int, category string, isParis bool) (bool, utils.Error)
	ListCostants() (domain.Costants, utils.Error)
	ListNamesakes(page, size int) ([]usecases.Namesake, utils.Error)
	ListPartecipations(page, size int) ([]domain.Participation, utils.Error)
	ListYears() ([]int, utils.Error)
	Login(user usecases.User) (string, utils.Error)
	ParseFileWithInfo(fileReader io.Reader, game domain.Game, format, city string) utils.Error
	PlayersCount() (int, utils.Error)
	PlayerHistory(player domain.Player) (usecases.SlimParticipationByYear, utils.Error)
	UpdateCostants(costants domain.Costants) utils.Error
	UpdateNamesake(n usecases.Namesake) utils.Error
	Log(msg string, args ...interface{})
}

// WebserviceHandler represents the mechanism that transform HTTP requests to
// data that the usecases layer can comprehend
type WebserviceHandler struct {
	Interactor Interactor
}
