package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CanobbioE/reelo/backend/services"
)

// GetRanks returns a list of all the ranks
// TODO: filters maybe
func GetRanks(w http.ResponseWriter, r *http.Request) {
	pages := r.URL.Query().Get("page")
	sizes := r.URL.Query().Get("size")

	page, err := strconv.Atoi(string(pages))
	if err != nil {
		log.Printf("Error converting page: %v", err)
		http.Error(w, "cannot parse query string", http.StatusBadRequest)
		return
	}

	size, err := strconv.Atoi(string(sizes))
	if err != nil {
		log.Printf("Error converting size: %v", err)
		http.Error(w, "cannot parse query string", http.StatusBadRequest)
		return
	}

	ranks, err := services.GetRanks(page, size)
	if err != nil {
		log.Printf("Error getting ranks: %v", err)
		http.Error(w, "cannot get ranks", http.StatusInternalServerError)
		return
	}

	for i, r := range ranks {
		history, err := services.GetHistory(r.Name, r.Surname)
		if err != nil {
			log.Printf("Error getting history: %v", err)
			http.Error(w, "cannot get history", http.StatusInternalServerError)
			return
		}
		ranks[i].History = history
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

// GetPlayersCount returns how many players are stored in the DB
func GetPlayersCount(w http.ResponseWriter, r *http.Request) {

	count, err := services.GetPlayersCount()
	if err != nil {
		log.Printf("Error getting count: %v", err)
		http.Error(w, "cannot get count", http.StatusBadRequest)
		return
	}

	ret, err := json.Marshal(count)
	if err != nil {
		log.Printf("Error marshalling count: %v", err)
		http.Error(w, "cannot marshal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return

}
