package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/services"
)

// GetYears returns a list of all the stored years
func GetYears(w http.ResponseWriter, r *http.Request) {

	years, err := services.GetYears()
	if err != nil {
		log.Printf("Error getting years: %v", err)
		http.Error(w, "cannot get years", http.StatusBadRequest)
		return
	}

	ret, err := json.Marshal(years)
	if err != nil {
		log.Printf("Error marshalling years: %v", err)
		http.Error(w, "cannot marshal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return

}
