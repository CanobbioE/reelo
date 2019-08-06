package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	// We want to take the error returned by the parser
	// and have it displayed in the FE
	err = services.ParseFileWithInfo(file, uploadInfo)
	if err != nil {
		log.Printf("Error while parsing file: %v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	log.Printf("\n\nFile parsed succesfully\n")

	go func() {
		err = services.CalculateAllReelo(true)
		if err != nil {
			log.Printf("Error recalculating reelo file: %v", err)
			return
		}
		log.Println("Recalculated REELO for all players")
	}()

	services.Backup()
	return
}
