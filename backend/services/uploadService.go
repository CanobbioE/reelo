package services

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/CanobbioE/reelo/backend/api"
	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/utils/parse"
)

// ParseFileWithInfo reads the content of a buffer and verifies its correctness
// and tries to parse each line into an entity to be saved in the database
// appending the data defined in the upload info
// TODO: implement this
func ParseFileWithInfo(fileReader io.Reader, info api.UploadInfo) error {
	db := rdb.NewDB()
	var results []parse.LineInfo
	year, err := strconv.Atoi(info.Year)
	if err != nil {
		return err
	}
	category := strings.ToUpper(info.Category)
	format := parse.NewFormat(strings.Split(info.Format, " "))
	isParis := info.IsParis

	// TODO: this is the warning we want to return to the front end
	results, warning := parse.File(fileReader, format, year, category)
	if warning != nil {
		log.Printf("parse.File returned warning: %v\n", warning)
		return warning
	}

	db.InserRankingFile(results, year, category, isParis)
	return nil
}
