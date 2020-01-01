package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/dto"
	"github.com/CanobbioE/reelo/backend/services"
)

// HandleAlgorithm handles which function get called based on the request method
func HandleAlgorithm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetCostants(w, r)
	case "PATCH":
		UpdateAlgorithm(w, r)
	}
	return
}

// UpdateAlgorithm updates some varaibles used for the reelo algorithm
func UpdateAlgorithm(w http.ResponseWriter, r *http.Request) {
	var c dto.Costants
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
	err = services.UpdateAlgorithm(context.Background(), c)
	if err != nil {
		log.Printf("Error updating costants: %v", err)
		http.Error(w, "can't update costants", http.StatusInternalServerError)
		return
	}
	log.Println("Costants updated")
	return
}

// GetCostants fetch the current values for the variables used
// in the reelo algorithm
func GetCostants(w http.ResponseWriter, r *http.Request) {
	vars, err := services.GetCostants()

	if err != nil {
		log.Printf("Error getting costants: %v", err)
		http.Error(w, "cannot get costants", http.StatusInternalServerError)
		return
	}

	ret, err := json.Marshal(vars)
	if err != nil {
		log.Printf("Error marshalling costants: %v", err)
		http.Error(w, "cannot marshal costants", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}
