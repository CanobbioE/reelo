package controllers

import (
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/services"
)

// ForceReelo forces the system to recalculate all the Reelo scores
func ForceReelo(w http.ResponseWriter, r *http.Request) {
	err := services.CalculateAllReelo(false)
	if err != nil {
		log.Printf("Error recalculating reelo: %v", err)
		return
	}
	log.Println("Recalculated (forced) Reelo for all players")

	return
}

// ForcePseudoReelo forces the system to recalculate all the Reelo
// and pseudo-Reelo scores
func ForcePseudoReelo(w http.ResponseWriter, r *http.Request) {
	err := services.CalculateAllReelo(true)
	if err != nil {
		log.Printf("Error recalculating reelo: %v", err)
		return
	}
	log.Println("Recalculated (forced) Reelo and pseudo-Reelo for all players")

	services.Backup()
	return
}
