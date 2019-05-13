package main

import (
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// TODO: on init check db integrity and recover data from files (data = formats, cities, parse.All)
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ranks", controllers.GetRanks).Methods("GET")
	router.HandleFunc("/admin", controllers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/upload", controllers.Upload).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}

// TODO: implement middleware
func requireAuth(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		log.Println("missing token")
		http.Error(w, "missing token", http.StatusUnauthorized)
	}
}
