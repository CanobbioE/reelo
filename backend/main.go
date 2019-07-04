package main

/* TODOs to ask to scientific committee
* Double check category order
* How to handle paris (do I exclude results from paris on every query?
 */
import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/CanobbioE/reelo/backend/controllers"
	"github.com/CanobbioE/reelo/backend/middlewares"
	"github.com/CanobbioE/reelo/backend/services"
	"github.com/CanobbioE/reelo/backend/utils/parse"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
)

func init() {
	// TODO: Check db integrity
	// TODO: Call parse.All() if we have stuff in Ranks folder
	parse.GetCities()
	log.Println("App initialized")
}

func main() {
	// Configure Logging
	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath != "" {
		f, err := os.OpenFile("./"+logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)
	}

	// Backup scheduling
	gocron.Every(1).Days().At("03:00").Do(services.Backup)
	gocron.Start()

	router := mux.NewRouter()

	router.HandleFunc("/ranks", controllers.GetRanks).Methods("GET")
	router.HandleFunc("/admin", controllers.Login).Methods("POST")
	router.HandleFunc("/upload", middlewares.Auth(
		http.HandlerFunc(controllers.Upload))).Methods("POST")
	router.HandleFunc("/force-reelo", middlewares.Auth(
		http.HandlerFunc(controllers.ForceReelo))).Methods("PUT")
	router.HandleFunc("/algorithm", middlewares.Auth(
		http.HandlerFunc(controllers.HandleAlgorithm))).Methods("PATCH", "GET")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}