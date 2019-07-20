package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/dto"
	"github.com/CanobbioE/reelo/backend/utils/parse"
)

// ParseFileWithInfo reads the content of a buffer and verifies its correctness
// and tries to parse each line into an entity to be saved in the database
// appending the data defined in the upload info
func ParseFileWithInfo(fileReader io.Reader, info dto.UploadInfo) error {
	db := rdb.NewDB()
	defer db.Close()

	var results []parse.LineInfo
	year, err := strconv.Atoi(info.Year)
	if err != nil {
		return err
	}
	start, err := strconv.Atoi(info.Start)
	if err != nil {
		return err
	}
	end, err := strconv.Atoi(info.End)
	if err != nil {
		return err
	}
	category := strings.ToUpper(info.Category)
	format, err := parse.NewFormat(strings.Split(info.Format, " "))
	if err != nil {
		return err
	}
	isParis := info.IsParis

	// this is the warning we want to return to the front end
	results, warning := parse.File(fileReader, format, year, category)
	if warning != nil {
		log.Printf("parse.File() returned warning: %v\n", warning)
		return warning
	}

	gameInfo := rdb.GameInfo{
		year,
		category,
		start,
		end,
	}

	db.InserRankingFile(context.Background(), results, gameInfo, isParis)
	return nil
}

// SaveRankingFile saves the specified reader in a file named "year_category.txt"
// TODO: this is unusued due to limited space on server
func SaveRankingFile(src io.Reader, year, category string, isParis bool) error {
	prefix := fmt.Sprintf("%s/%s", parse.RankPath, year)
	if isParis {
		prefix = fmt.Sprintf("%s/paris", prefix)
	}
	err := os.MkdirAll(prefix, 0777)
	if err != nil && os.IsNotExist(err) {
		return err
	}

	dstPath := fmt.Sprintf("%s/%s_%s.txt", prefix, year, strings.ToUpper(category))
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil

}
