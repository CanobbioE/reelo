package webinterface

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CanobbioE/reelo/backend/interfaces/webinterface/dto"
	"github.com/CanobbioE/reelo/backend/utils"
)

// ListRanks returns a list of all the ranks
func (wh *WebserviceHandler) ListRanks(w http.ResponseWriter, r *http.Request) {
	page, size, err := utils.Paginate(r)
	if !err.IsNil {
		wh.Interactor.Log("ListRanks: error paginating: %v", err.String())
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}

	var ranks []dto.Rank
	partecipations, err := wh.Interactor.ListRanks(page, size)
	if !err.IsNil {
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}

	for _, p := range partecipations {
		history, err := wh.Interactor.PlayerHistory(p.Player)
		if !err.IsNil {
			http.Error(w, err.String(), err.HTTPStatus)
			return
		}
		ranks = append(ranks, dto.Rank{
			Player:       p.Player,
			History:      history,
			LastCategory: p.Game.Category,
		})
	}

	ret, e := json.Marshal(ranks)
	if e != nil {
		wh.Interactor.Log("ListRanks: cannot marshal ranks: %v", err)
		http.Error(w, utils.NewError(e, "E_GENERIC", http.StatusInternalServerError).String(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)

	return

}

// ListYears returns a list of all the stored years
func (wh *WebserviceHandler) ListYears(w http.ResponseWriter, r *http.Request) {
	years, err := wh.Interactor.ListYears()
	if !err.IsNil {
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}

	ret, e := json.Marshal(years)
	if e != nil {
		wh.Interactor.Log("ListYears: cannot marshal years: %v", e)
		http.Error(w, utils.NewError(e, "E_GENERIC", http.StatusInternalServerError).String(), http.StatusInternalServerError)
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
		wh.Interactor.Log("Upload: failed to receive the file: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	defer func() {
		if r := recover(); r != nil {
			wh.Interactor.Log("Upload: recovered: %v\n", r)
			err := fmt.Errorf("file corrupted: %v", r)
			http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
			return
		}
	}()

	var uploadInfo dto.FileUpload
	err = json.Unmarshal([]byte(r.FormValue("data")), &uploadInfo)
	if err != nil {
		wh.Interactor.Log("Upload: cannot unmarshal upload data: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}

	if err := wh.Interactor.DeleteIfAlreadyExists(uploadInfo.Game); !err.IsNil {
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}

	// We want to take the error returned by the parser
	// and have it displayed in the FE
	// TODO: error parser
	e := wh.Interactor.ParseFileWithInfo(file, uploadInfo.Game, uploadInfo.Format, uploadInfo.City)
	if !e.IsNil {
		http.Error(w, e.String(), e.HTTPStatus)
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
		wh.Interactor.Log("RankExistence: cannot read query string: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}
	category := string(r.URL.Query().Get("cat"))
	isParis, err := strconv.ParseBool(r.URL.Query().Get("isparis"))
	if err != nil {
		wh.Interactor.Log("RankExistence: cannot read query string: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}

	exists, e := wh.Interactor.DoesRankExist(year, category, isParis)
	if !e.IsNil {
		http.Error(w, e.String(), e.HTTPStatus)
		return
	}

	ret, err := json.Marshal(exists)
	if err != nil {
		wh.Interactor.Log("RankExistence: cannot marshal rank existence: %v", err)
		http.Error(w, utils.NewError(err, "E_GENERIC", http.StatusInternalServerError).String(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	return
}
