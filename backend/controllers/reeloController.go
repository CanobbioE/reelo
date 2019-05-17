package controllers

import (
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/services"
)

// ForceReelo forces the system to recalculate all the reelo scores
func ForceReelo(w http.ResponseWriter, r *http.Request) {
	services.CalculateAllReelo()
	log.Println("Recalculated (forced) REELO for all players")
	return
}
