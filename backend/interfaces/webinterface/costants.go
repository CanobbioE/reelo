package webinterface

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
)

// ListCostants fetch the current values used in the Reelo's algorithm.
func (wh *WebserviceHandler) ListCostants(w http.ResponseWriter, r *http.Request) {
	constants, e := wh.Interactor.ListCostants()
	if !e.IsNil {
		http.Error(w, e.String(), http.StatusInternalServerError)
		return
	}

	ret, err := json.Marshal(constants)
	if err != nil {
		wh.Interactor.Log("ListCostants: cannot marshal costants: %v", err)
		http.Error(w, utils.NewError(err, "E_GENERIC", http.StatusInternalServerError).String(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}

// UpdateCostants updates some of the values used to calculate the players Reelo.
func (wh *WebserviceHandler) UpdateCostants(w http.ResponseWriter, r *http.Request) {
	var c domain.Costants
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wh.Interactor.Log("Error while reading costants body: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		wh.Interactor.Log("Error while unmarshalling costants: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}
	e := wh.Interactor.UpdateCostants(c)
	if !e.IsNil {
		wh.Interactor.Log("Error updating costants: %v", e.Message)
		http.Error(w, e.String(), e.HTTPStatus)
		return
	}
	wh.Interactor.Log("Costants updated")
	return
}
