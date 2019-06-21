package services

import (
	"log"

	rdb "github.com/CanobbioE/reelo/backend/db"
)

// Backup performs the backup of the database
func Backup() {
	log.Println("Backup started")
	db := rdb.NewDB()
	defer db.Close()
	file := db.Backup()
	err := UploadFile(file)
	if err != nil {
		log.Printf("Error while uploading file to google drive: %v\n", err)
		return
	}
	log.Println("Backup successfully done!")
}
