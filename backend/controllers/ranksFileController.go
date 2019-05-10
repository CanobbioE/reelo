package controllers

import (
	"encoding/json"
	"net/http"
)

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

// Stuff does things
func Stuff(w http.ResponseWriter, r *http.Request) {
	// db := rdb.NewDB()
	// dataAll := parse.All()
	// for year, lines := range dataAll {
	//    for _, line := range lines {
	//       playerID := db.Add(context.Background(), "giocatore", line.name, line.surname, 0)
	//       resultID := db.Add(context.Background(), "risultato", line.tempo, line.esercizi, line.punteggio)
	//       gameID := db.Add(context.Background(), "giochi", year, line.categoria)
	//       db.Add(context.Background(), "partecipazione", playerID, gameID, resultID, line.sede)

	//    }
	// }

	// TODO: import to db
}
