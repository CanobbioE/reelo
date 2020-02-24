package webinterface

import (
	"fmt"
	"io"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
)

// Interactor define the usecases behaviour, using an interface allows
// to change implementation as needed
type Interactor interface {
	AnalysisHistory(player domain.Player) (usecases.HistoryByYear, []int, error)
	AddComment(namesake usecases.Namesake) error
	CalculateAllReelo(doPseudo bool) error
	CalculateMaxScoreObtainable(game domain.Game) (int, error)
	CalculatePlayerReelo(player domain.Player, doPseudo bool) error
	DeleteIfAlreadyExists(game domain.Game) error
	DoesRankExist(year int, category string, isParis bool) (bool, error)
	ListCostants() (domain.Costants, error)
	ListNamesakes(page, size int) ([]usecases.Namesake, error)
	ListRanks(page, size int) ([]domain.Partecipation, error)
	ListYears() ([]int, error)
	Login(user usecases.User) (string, fmt.Stringer)
	ParseFileWithInfo(fileReader io.Reader, game domain.Game, format, city string) error
	PlayersCount() (int, error)
	PlayerHistory(player domain.Player) (usecases.SlimPartecipationByYear, error)
	UpdateCostants(costants domain.Costants) error
	UpdateNamesake(n usecases.Namesake) error
	Log(msg string, args ...interface{})
	Error(err error, code string, httpStatus int) string
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
