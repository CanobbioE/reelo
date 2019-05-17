package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/services"
)

// GetRanks returns a list of all the ranks
// TODO: filters maybe
// TODO: pagination I guess
func GetRanks(w http.ResponseWriter, r *http.Request) {
	ranks, err := services.GetRanks()
	if err != nil {
		log.Printf("Error getting ranks: %v", err)
		http.Error(w, "cannot get ranks", http.StatusInternalServerError)
		return
	}

	ret, err := json.Marshal(ranks)
	if err != nil {
		log.Printf("Error marshalling ranks: %v", err)
		http.Error(w, "cannot marshal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return

}
