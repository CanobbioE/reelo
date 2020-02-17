package webinterface

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/usecases"
)

// PlayersCount returns how many players are stored in the DB
func (wh *WebserviceHandler) PlayersCount(w http.ResponseWriter, r *http.Request) {
	count, err := wh.Interactor.PlayersCount()
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

// ForcePseudoReelo does stuff
func (wh *WebserviceHandler) ForcePseudoReelo(w http.ResponseWriter, r *http.Request) {
	err := wh.Interactor.CalculateAllReelo(true)
	if err != nil {
		log.Printf("Error recalculating reelo: %v", err)
		return
	}
	log.Println("Recalculated (forced) Reelo and pseudo-Reelo for all players")

	// TODO wh.Interactor.Backup()
	return
}

// AddComment adds a comment to a namesake
func (wh *WebserviceHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	var n usecases.Namesake
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
	if err = wh.Interactor.AddComment(n); err != nil {
		log.Printf("Error commenting namesake: %v", err)
		http.Error(w, "can't comment namesake", http.StatusInternalServerError)
		return
	}
	return
}
