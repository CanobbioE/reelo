package webinterface

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
)

// PlayersCount returns the number of player the application is storing.
func (wh *WebserviceHandler) PlayersCount(w http.ResponseWriter, r *http.Request) {
	count, err := wh.Interactor.PlayersCount()
	if !err.IsNil {
		http.Error(w, err.String(), http.StatusBadRequest)
		return
	}

	ret, e := json.Marshal(count)
	if e != nil {
		wh.Interactor.Log("PlayersCount: cannot marshal count: %v", e)
		http.Error(w, utils.NewError(e, "E_GENERIC", http.StatusInternalServerError).String(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}

// CalculateReelo forces the system to recalculate all the "pseudo reelo"
// and reelo scores
func (wh *WebserviceHandler) CalculateReelo(w http.ResponseWriter, r *http.Request) {
	err := wh.Interactor.CalculateAllReelo()
	if !err.IsNil {
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}
	wh.Interactor.Log("Recalculated (forced) Reelo and pseudo-Reelo for all players")

	return
}

// AddComment adds a comment to a player.
func (wh *WebserviceHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	var c domain.Comment
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wh.Interactor.Log("AddComment: cannot read body: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		wh.Interactor.Log("AddComment: cannot unmarshal body: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}
	if e := wh.Interactor.AddComment(c); !e.IsNil {
		http.Error(w, e.String(), e.HTTPStatus)
		return
	}
	return
}
