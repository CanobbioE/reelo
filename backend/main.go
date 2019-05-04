package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/parse"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/upload", CreateRankingFile).Methods("POST")
	router.HandleFunc("/ranks", GetRanks).Methods("GET")
	router.HandleFunc("/stuff", Stuff).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Stuff does things
func Stuff(w http.ResponseWriter, r *http.Request) {
	db := rdb.NewDB()
	dataAll := parse.All()
	for year, lines := range dataAll {
		for _, line := range lines {
			playerID := db.Add(context.Background(), "giocatore", line.name, line.surname, 0)
			resultID := db.Add(context.Background(), "risultato", line.tempo, line.esercizi, line.punteggio)
			gameID := db.Add(context.Background(), "giochi", year, line.categoria)
			db.Add(context.Background(), "partecipazione", playerID, gameID, resultID, line.sede)

		}
	}

	// TODO: import to db
}

// GetRanks returns a list of all the ranks in the database
func GetRanks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
}

// CreateRankingFile creates a new ranking file
// TODO: authentication
func CreateRankingFile(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// var person Person
	// _ = json.NewDecoder(r.Body).Decode(&person)
	// person.ID = params["id"]
	// people = append(people, person)
	// json.NewEncoder(w).Encode(people)
}
