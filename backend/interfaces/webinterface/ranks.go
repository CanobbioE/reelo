package webinterface

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils"
)

// ListRanks returns a list of all the ranks
func (wh *WebserviceHandler) ListRanks(w http.ResponseWriter, r *http.Request) {
	page, size, err := utils.Paginate(r)
	if err != nil {
		log.Printf("Error paginating ranks: %v", err)
		http.Error(w, "cannot parse query string", http.StatusBadRequest)
		return
	}

	_, err = wh.Interactor.ListRanks(page, size)
	if err != nil {
		log.Printf("Error getting ranks: %v", err)
		http.Error(w, "cannot get ranks", http.StatusInternalServerError)
		return
	}
	/* TODO
	for i, p := range partecipations {
		history, err := wh.Interactor.PlayerHistory(p.Player)
		if err != nil {
			log.Printf("Error getting history: %v", err)
			http.Error(w, "cannot get history", http.StatusInternalServerError)
			return
		}
		partecipations[i].History = history
	}

	ret, err := json.Marshal(ranks)
	if err != nil {
		log.Printf("Error marshalling ranks: %v", err)
		http.Error(w, "cannot marshal", http.StatusInternalServerError)
		return
	}
	*/
	w.Header().Set("Content-Type", "application/json")
	w.Write(nil)
	return

}

// ListYears returns a list of all the stored years
func (wh *WebserviceHandler) ListYears(w http.ResponseWriter, r *http.Request) {
	years, err := wh.Interactor.ListYears()
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

// Upload creates a new ranking file
func (wh *WebserviceHandler) Upload(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error receiving the file: %v", err)
		http.Error(w, "invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in Upload: %v\n", r)
			http.Error(w, fmt.Sprintf("file corrupted: %v", r), http.StatusBadRequest)
			return
		}
	}()

	var game domain.Game
	err = json.Unmarshal([]byte(r.FormValue("data")), &uploadInfo)
	if err != nil {
		log.Printf("Error while unmarshalling upload data: %v", err)
		http.Error(w, "can't unmarshal data", http.StatusBadRequest)
		return
	}

	if err := wh.Interactor.DeleteIfAlreadyExists(game); err != nil {
		log.Printf("Error while checking ranks existence: %v", err)
		http.Error(w, "can't check existence", http.StatusBadRequest)
		return
	}

	// We want to take the error returned by the parser
	// and have it displayed in the FE
	// TODO: error parser
	err = wh.Interactor.ParseFileWithInfo(file, game)
	if err != nil {
		log.Printf("Error while parsing file: %v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	// TODO wh.Interactor.Backup()
	return
}

// RankExistence is called to verify if a year-category ranking file has been already uploaded
func (wh *WebserviceHandler) RankExistence(w http.ResponseWriter, r *http.Request) {
	yearString := string(r.URL.Query().Get("y"))
	year, err := strconv.Atoi(yearString)
	if err != nil {
		log.Printf("Error reading query string: %v", err)
		http.Error(w, fmt.Sprintf("can't read query string: %v", err), http.StatusInternalServerError)
		return
	}
	category := string(r.URL.Query().Get("cat"))
	isParis, err := strconv.ParseBool(r.URL.Query().Get("isparis"))
	if err != nil {
		log.Printf("Error reading query string: %v", err)
		http.Error(w, fmt.Sprintf("can't read query string: %v", err), http.StatusInternalServerError)
		return
	}

	exists, err := wh.Interactor.DoesRankExist(year, category, isParis)
	if err != nil {
		log.Printf("Error checking rank existence: %v", err)
		http.Error(w, fmt.Sprintf("can't check existencce: %v", err), http.StatusInternalServerError)
		return
	}

	ret, err := json.Marshal(exists)
	if err != nil {
		log.Printf("Error marshalling exists: %v", err)
		http.Error(w, "cannot marshal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}
