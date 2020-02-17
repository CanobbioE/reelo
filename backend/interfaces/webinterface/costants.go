package webinterface

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/domain"
)

// ListCostants fetch the current values for the constants used in the reelo algorithm
func (wh *WebserviceHandler) ListCostants(w http.ResponseWriter, r *http.Request) {
	constants, err := wh.Interactor.ListCostants()

	if err != nil {
		log.Printf("Error getting costants: %v", err)
		http.Error(w, "cannot get costants", http.StatusInternalServerError)
		return
	}

	ret, err := json.Marshal(constants)
	if err != nil {
		log.Printf("Error marshalling costants: %v", err)
		http.Error(w, "cannot marshal costants", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}

// UpdateCostants updates some constants used for the reelo algorithm
func (wh *WebserviceHandler) UpdateCostants(w http.ResponseWriter, r *http.Request) {
	var c domain.Costants
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error while  reading costants body: %v", err)
		http.Error(w, "can't update costants", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Printf("Error while unmarshalling costants: %v", err)
		http.Error(w, "can't update costants", http.StatusBadRequest)
		return
	}
	err = wh.Interactor.UpdateCostants(c)
	if err != nil {
		log.Printf("Error updating costants: %v", err)
		http.Error(w, "can't update costants", http.StatusInternalServerError)
		return
	}
	log.Println("Costants updated")
	return
}
