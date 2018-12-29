package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var FolderPath = "./ranks"

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

func indexOf(ss []string, pattern string) int {
	for i, s := range ss {
		if s == pattern {
			return i
		}
	}
	return -1
}

func updateDb(data []string, f *Format) {
	// if !db.contains("giocatore", nome, cognome)
	// 	db.add("giocatore", nome, cognome)
	// 	db.add("risultato", eserciizi, punteggio, tempo)
	// 	ottieni riferimento a GIOCHI trmite anno e categoria
	// 	db.add("partecipazione", giocatore, risultato, giochi)
}
