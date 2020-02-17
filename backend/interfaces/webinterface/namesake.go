package webinterface

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// ListNamesakes identifies and returns a list of all the namesakes
// Paging is available and optional, beware that the page size is not used for
// the number of namesakes that we want to return but for how many players
// will be players analyzed.
func (wh *WebserviceHandler) ListNamesakes(w http.ResponseWriter, r *http.Request) {
	page, size, err := utils.Paginate(r)
	if err != nil {
		log.Printf("Error paginating namesakes: %v", err)
		http.Error(w, "cannot parse query string", http.StatusBadRequest)
		return
	}

	namesakes, err := wh.Interactor.ListNamesakes(page, size)
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
func (wh *WebserviceHandler) UpdateNamesake(w http.ResponseWriter, r *http.Request) {
	var n usecases.Namesake
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

	err = wh.Interactor.UpdateNamesake(n)
	if err != nil {
		log.Printf("Error updating namesake: %v", err)
		http.Error(w, "can't update namesake", http.StatusInternalServerError)
		return
	}
}
