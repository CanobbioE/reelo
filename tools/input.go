package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	rdb "github.com/CanobbioE/reelo/db"
)

var FolderPath = "./ranks"
var db = rdb.DB()

// parseRankingFile reads a ranking from the correct file using the specified
// format. The file's name must be "year_category.txt"
func parseRankingFile(year int, category, format string) {
	filePath := fmt.Sprintf("%s/%d/%d_%s.txt", FolderPath, year, year, category)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f := newFormat(strings.Split(strings.ToLower(format), " "))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		updateDb(data, f)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

// Format represents the format used in a ranking file.
// Each field represents an information's index inside a line of the file.
type Format struct {
	Exercises, Score, Time int
	Name, Surname          int
	City                   int
}

// newFormat returns a new format based on the slice of string passed
func newFormat(ff []string) *Format {
	return &Format{
		Name:      indexOf(ff, "nome"),
		Surname:   indexOf(ff, "cognome"),
		City:      indexOf(ff, "sede"),
		Exercises: indexOf(ff, "esercizi"),
		Score:     indexOf(ff, "punteggio"),
		Time:      indexOf(ff, "tempo"),
	}
}

// indexOf returns the position of the pattern's first occurency inside the ss slice
func indexOf(ss []string, pattern string) int {
	for i, s := range ss {
		if s == pattern {
			return i
		}
	}
	return -1
}

// updateDb updates the database by adding all the data read from a single row
// of the Ranking's file
func updateDb(data []string, f *Format) {
	var time = "0"
	var pId int
	// Add the player information only if he doesn't already exist in the db
	if !rdb.ContainsPlayer(db, data[f.Name], data[f.Surname]) {
		pId = rdb.Add(db, "giocatore", data[f.Name], data[f.Surname])
	} else {
		pId = 0 // select id from giocatore where nome = cognome and cognome = surname
	}
	if f.Time != -1 {
		time = data[f.Time]
	}
	rId := rdb.Add(db, "risultato", data[f.Exercises], data[f.Score], time)
	// TODO ottieni riferimento a GIOCHI trmite anno e categoria
	rdb.Add(db, "partecipazione", pId, rId, "giochiID")
}
