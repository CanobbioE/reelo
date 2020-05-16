package webinterface

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// ListNamesakes identifies and returns a list of all the namesakes
// that cannot be automatically solved.
// Paging is available and optional.
// Page and size indicate how many players will be analyzed.
func (wh *WebserviceHandler) ListNamesakes(w http.ResponseWriter, r *http.Request) {
	page, size, err := utils.Paginate(r)
	if !err.IsNil {
		wh.Interactor.Log("ListNamesakes: error paginating: %v", err.String())
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}

	namesakes, err := wh.Interactor.ListNamesakes(page, size)
	if !err.IsNil {
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}

	ret, e := json.Marshal(namesakes)
	if e != nil {
		wh.Interactor.Log("ListNamesakes: cannot marshal namesakes: %v", e)
		http.Error(w, utils.NewError(e, "E_GENERIC", http.StatusInternalServerError).String(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}

// UpdateNamesake changes the history for a given player
// TODO: soon obsolete
func (wh *WebserviceHandler) UpdateNamesake(w http.ResponseWriter, r *http.Request) {
	var n usecases.Namesake
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wh.Interactor.Log("UpdateNamesake: cannot read body: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &n)
	if err != nil {
		wh.Interactor.Log("UpdateNamesake: cannot unmarshal body: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}

	e := wh.Interactor.UpdateNamesake(n)
	if !e.IsNil {
		http.Error(w, e.String(), e.HTTPStatus)
		return
	}
}
