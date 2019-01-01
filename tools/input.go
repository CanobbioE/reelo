package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	rdb "github.com/CanobbioE/reelo/db"
)

const RANK_PATH = "./ranks"

// parseRankingFile reads a ranking from the correct file using the specified
// format. The file's name must be "year_category.txt"
func parseRankingFile(year int, category, format string) {
	filePath := fmt.Sprintf("%s/%d/%d_%s.txt", RANK_PATH, year, year, category)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// add the current year+category to the db
	db := rdb.DB()
	gId := rdb.Add(db, "giochi", year, category)
	db.Close()

	f := newFormat(strings.Split(strings.ToLower(format), " "))

	// Here we do stuff and things with the input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		updateDb(data, f, gId)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

// updateDb updates the database by adding all the data read from a single row
// of the Ranking's file.
func updateDb(data []string, f *Format, gId int) {
	var pId, time int
	db := rdb.DB()
	defer db.Close()
	// Add the player information only if he doesn't already exist in the db
	if !rdb.ContainsPlayer(db, data[f.Name], data[f.Surname]) {
		pId = rdb.Add(db, "giocatore", data[f.Name], data[f.Surname])
	} else {
		q := `
		SELECT id FROM Giocatore
		WHERE nome = ? AND cognome = ?
		`
		err := db.QueryRow(q, data[f.Name], data[f.Surname]).Scan(&pId)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Sometime the time is not specified in the ranking file
	if f.Time != -1 {
		var err error
		time, err = strconv.Atoi(data[f.Time])
		if err != nil {
			time = 0
		}
	}
	rId := rdb.Add(db, "risultato", time, data[f.Exercises], data[f.Score])
	rdb.Add(db, "partecipazione", pId, gId, rId)
}
