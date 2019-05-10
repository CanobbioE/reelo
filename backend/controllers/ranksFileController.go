package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/api"
	"github.com/CanobbioE/reelo/backend/services"
)

// GetRanks returns a list of all the ranks in the database
func GetRanks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
}

// Upload creates a new ranking file
// TODO: authentication
func Upload(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	// TODO: parse file, return errors so that user can correct them
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error receiving the file: %v", err)
		http.Error(w, "invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)

	var uploadInfo api.UploadInfo
	err = json.Unmarshal([]byte(r.FormValue("data")), &uploadInfo)
	if err != nil {
		log.Printf("Error while unmarshalling data: %v", err)
		http.Error(w, "can't unmarshal data", http.StatusBadRequest)
		return
	}

	// TODO here we should take the contents of the file,
	// parse it and then save it to db.
	// We also want to take the error returned by the parser
	// and have it displayed in the FE
	err = services.ParseFileWithInfo(buf, uploadInfo)
	if err != nil {
		log.Printf("Error while parsing file: %v", err)
		http.Error(w, "TODO error", http.StatusInternalServerError)
		return
	}

	return
}
