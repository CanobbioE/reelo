package controllers

import (
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/services"
)

// PurgePlayers fixes all possible mistakes suche as namesakes
func PurgePlayers(w http.ResponseWriter, r *http.Request) {
	if err := services.PurgeNamesakes(); err != nil {
		log.Println(err)
		http.Error(w, "cannot purge namesakes", http.StatusBadRequest)
		return
	}
	log.Println("Purged all players")
	return
}
