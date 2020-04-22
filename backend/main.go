package main

/* TODO:
 * change elo service to use only IDs and not name/surname
 */
import (
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/infrastructure"
	"github.com/CanobbioE/reelo/backend/infrastructure/mysqlhandler"
	"github.com/CanobbioE/reelo/backend/interfaces/repository"
	"github.com/CanobbioE/reelo/backend/interfaces/webinterface"
	mw "github.com/CanobbioE/reelo/backend/interfaces/webinterface/middlewares"
	"github.com/CanobbioE/reelo/backend/pkg/parse"
	"github.com/CanobbioE/reelo/backend/usecases/interactor"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	i  *interactor.Interactor
	wh webinterface.WebserviceHandler
)

func init() {
	parse.GetCities()
	log.Println("Creating logger...")
	logger := infrastructure.NewLogger()

	log.Println("Connecting to the database...")
	dbHandler, err := mysqlhandler.InitRepoHandler()
	if err != nil {
		log.Fatalf("Cannot instantiate repository handler: %v", err)
	}
	dbHandlers := make(map[string]repository.DbHandler)
	for _, repo := range repository.All() {
		dbHandlers[repo] = dbHandler
	}

	i = &interactor.Interactor{
		CommentRepository:       repository.NewDbCommentRepo(dbHandlers),
		CostantsRepository:      repository.NewDbCostantsRepo(dbHandlers),
		GameRepository:          repository.NewDbGameRepo(dbHandlers),
		ParticipationRepository: repository.NewDbParticipationRepo(dbHandlers),
		PlayerRepository:        repository.NewDbPlayerRepo(dbHandlers),
		ResultRepository:        repository.NewDbResultRepo(dbHandlers),
		UserRepository:          repository.NewDbUserRepo(dbHandlers),
		HistoryRepository:       repository.NewDbHistoryRepo(dbHandlers),
		Logger:                  logger,
	}
	wh = webinterface.WebserviceHandler{Interactor: i}
	log.Println("Initializing reelo algorithm...")
	i.InitCostants()
	log.Println("App initialized!")
}

func main() {
	// Backup scheduling
	// gocron.Every(1).Days().At("03:00").Do(services.Backup)
	// gocron.Start()

	router := mux.NewRouter()

	// Routing
	// endpoint /players
	router.HandleFunc("/players/count", wh.PlayersCount).Methods("GET")
	router.HandleFunc("/players/reelo/calculate", mw.RequireAuth(
		http.HandlerFunc(wh.ForcePseudoReelo))).Methods("POST", "PUT")
	router.HandleFunc("/players/comment", mw.RequireAuth(
		http.HandlerFunc(wh.AddComment))).Methods("POST")

	// endpoint /ranks
	router.HandleFunc("/ranks/all/", wh.ListRanks).Methods("GET")
	router.HandleFunc("/ranks/upload", mw.RequireAuth(
		http.HandlerFunc(wh.Upload))).Methods("POST")
	router.HandleFunc("/ranks/exist/", wh.RankExistence).Methods("GET")
	router.HandleFunc("/ranks/years", wh.ListYears).Methods("GET")

	// endpoint /auth
	router.HandleFunc("/auth/login", wh.Login).Methods("POST")

	// endpoint /namesakes
	router.HandleFunc("/namesakes/all", mw.RequireAuth(
		http.HandlerFunc(wh.ListNamesakes))).Methods("GET")
	router.HandleFunc("/namesakes/update", mw.RequireAuth(
		http.HandlerFunc(wh.UpdateNamesake))).Methods("POST")

	// endpoint /costants
	router.HandleFunc("/costants/all", mw.RequireAuth(
		http.HandlerFunc(wh.ListCostants))).Methods("GET")
	router.HandleFunc("/costants/update", mw.RequireAuth(
		http.HandlerFunc(wh.UpdateCostants))).Methods("POST", "PATCH")

	// Serving
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
