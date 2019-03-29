package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	// TODO: import to db
	fmt.Println(parse.All())
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
