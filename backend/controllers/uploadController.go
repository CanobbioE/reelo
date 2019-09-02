package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CanobbioE/reelo/backend/dto"
	"github.com/CanobbioE/reelo/backend/services"
)

// Upload creates a new ranking file
func Upload(w http.ResponseWriter, r *http.Request) {
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

	var uploadInfo dto.UploadInfo
	err = json.Unmarshal([]byte(r.FormValue("data")), &uploadInfo)
	if err != nil {
		log.Printf("Error while unmarshalling upload data: %v", err)
		http.Error(w, "can't unmarshal data", http.StatusBadRequest)
		return
	}

	if err := services.DeleteIfAlreadyExists(uploadInfo); err != nil {
		log.Printf("Error while checking ranks existence: %v", err)
		http.Error(w, "can't check existence", http.StatusBadRequest)
		return
	}

	// We want to take the error returned by the parser
	// and have it displayed in the FE
	// TODO: error parser
	err = services.ParseFileWithInfo(file, uploadInfo)
	if err != nil {
		log.Printf("Error while parsing file: %v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	services.Backup()
	return
}

// CheckRankExistence is called to verify if a year-category ranking file has been already uploaded
func CheckRankExistence(w http.ResponseWriter, r *http.Request) {
	year := string(r.URL.Query().Get("y"))
	category := string(r.URL.Query().Get("cat"))
	isParis, err := strconv.ParseBool(r.URL.Query().Get("isparis"))
	if err != nil {
		log.Printf("Error reading query string: %v", err)
		http.Error(w, fmt.Sprintf("can't read query string: %v", err), http.StatusInternalServerError)
		return
	}

	exists, err := services.DoesRankExist(year, category, isParis)
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
