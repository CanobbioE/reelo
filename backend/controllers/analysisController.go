package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/services"
	"github.com/CanobbioE/reelo/backend/utils"
)

// GetNamesakes identifies and returns a list of all the namesakes
// Paging is available and optional, beware that the page size is not used for
// the number of namesakes that we want to return but for how many players
// will be players analyzed.
func GetNamesakes(w http.ResponseWriter, r *http.Request) {
	page, size, err := utils.Paginate(r)
	if err != nil {
		log.Printf("Error paginating namesakes: %v", err)
		http.Error(w, "cannot parse query string", http.StatusBadRequest)
		return
	}

	namesakes, err := services.GetNamesakes(page, size)
	if err != nil {
		log.Println(err)
		http.Error(w, "cannot get namesakes", http.StatusBadRequest)
		return
	}
	log.Println("Done getting!")

	ret, err := json.Marshal(namesakes)
	if err != nil {
		log.Printf("Error marshalling namesakes: %v", err)
		http.Error(w, "cannot marshal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}
