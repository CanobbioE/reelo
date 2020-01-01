package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/dto"
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

// UpdateNamesake changes the history for a given player
func UpdateNamesake(w http.ResponseWriter, r *http.Request) {
	var n dto.Namesake
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error while reading namesake body: %v", err)
		http.Error(w, "can't update namesake", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Printf("Error while unmarshalling namesake: %v", err)
		http.Error(w, "can't update namesake", http.StatusBadRequest)
		return
	}

	err = services.UpdateNamesake(n)
	if err != nil {
		log.Printf("Error updating namesake: %v", err)
		http.Error(w, "can't update namesake", http.StatusInternalServerError)
		return
	}
}

// CommentNamesake adds a comment to a namesake
func CommentNamesake(w http.ResponseWriter, r *http.Request) {
	var n dto.Namesake
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error while reading namesake body: %v", err)
		http.Error(w, "can't comment namesake", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Printf("Error while unmarshalling namesake: %v", err)
		http.Error(w, "can't comment namesake", http.StatusBadRequest)
		return
	}
	if err = services.CommentNamesake(n); err != nil {
		log.Printf("Error commenting namesake: %v", err)
		http.Error(w, "can't comment namesake", http.StatusInternalServerError)
		return
	}
	return
}
