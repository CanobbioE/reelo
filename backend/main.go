package main

import (
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ranks", controllers.GetRanks).Methods("GET")
	router.HandleFunc("/admin", controllers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/upload", controllers.CreateRankingFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
