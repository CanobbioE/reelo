package webinterface

import (
	"io"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// Interactor define the usecases behaviour, using an interface allows
// to change implementation as needed
type Interactor interface {
	AnalysisHistory(player domain.Player) (usecases.HistoryByYear, []int, utils.Error)
	AddComment(c domain.Comment) utils.Error
	CalculateAllReelo(doPseudo bool) utils.Error
	CalculateMaxScoreObtainable(game domain.Game) (int, utils.Error)
	CalculatePlayerReelo(player domain.Player, doPseudo bool) utils.Error
	DeleteIfAlreadyExists(game domain.Game) utils.Error
	DoesRankExist(year int, category string, isParis bool) (bool, utils.Error)
	ListCostants() (domain.Costants, utils.Error)
	ListNamesakes(page, size int) ([]usecases.Namesake, utils.Error)
	ListRanks(page, size int) ([]domain.Partecipation, utils.Error)
	ListYears() ([]int, utils.Error)
	Login(user usecases.User) (string, utils.Error)
	ParseFileWithInfo(fileReader io.Reader, game domain.Game, format, city string) utils.Error
	PlayersCount() (int, utils.Error)
	PlayerHistory(player domain.Player) (usecases.SlimPartecipationByYear, utils.Error)
	UpdateCostants(costants domain.Costants) utils.Error
	UpdateNamesake(n usecases.Namesake) utils.Error
	Log(msg string, args ...interface{})
}

// WebserviceHandler represents the mechanism that transform HTTP requests to
// data that the usecases layer can comprehend
type WebserviceHandler struct {
	Interactor Interactor
}

/*
// UpdateDB is to be called from CLI, it is used to automate db updates.
// In production is an empty function
func UpdateDB(w http.ResponseWriter, r *http.Request) {
	log.Println("Called")
	ctx := context.Background()
	db := rdb.Instance()
	ids, err := db.AllPlayersID(ctx)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, id := range ids {
		history, years, err := db.AnalysisHistoryByID(ctx, id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accent := rdb.CreateAccent(history[years[0]][0].Year, 0, history[years[0]][0].City)
		if err = db.UpdateDB(ctx, accent, id); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	return
}
*/
