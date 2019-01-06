package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	rdb "github.com/CanobbioE/reelo/db"
)

const RANK_PATH = "../ranks"

var db *rdb.DB

/*
func init() {
	db = rdb.NewDB()
}
*/

type dataLine struct {
	Name      string
	Surname   string
	City      string
	Exercises int
	Points    int
	Time      int
}

// parseRankingFile reads a ranking from the correct file using the specified
// format. The file's name must be in the format of "year_category.txt"
func readRankingFile(year int, category string, format Format) {
	filePath := fmt.Sprintf("%s/%d/%d_%s.txt", RANK_PATH, year, year, category)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Couldn't open file.", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO: commenting this since we need to rewrite almost all of this.
		//data := strings.Split(scanner.Text(), " ")
		//updateDb(ctx, data, format, gID)
		parseLine(format, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLine(format Format, input string) {
	splitted := strings.Split(input, " ")
	var result dataLine

	fmt.Println(input)

	// Handling base case.
	// Otherwise it means there are fields with two or more words.
	if len(splitted) == len(format) {
		for fName, index := range format {
			var err error
			value := strings.Title(strings.ToLower(splitted[index]))
			switch fName {
			case "cognome":
				result.Surname = value
			case "nome":
				result.Name = value
			case "esercizi":
				result.Exercises, err = strconv.Atoi(value)
			case "punti":
				result.Points, err = strconv.Atoi(value)
			case "tempo":
				result.Time, err = strconv.Atoi(value)
			case "città":
				result.City = strings.Title(value)
			case "città(provincia)":
				// TODO delete provincia from the input
				result.City = strings.Title(value)
			}
			if err != nil {
				log.Fatal("Could not convert data.", err)
			}
		}
		fmt.Println(result)
	} else {
		log.Print("Not implemented")
	}
	fmt.Println()
}

/*
// updateDb updates the database by adding all the data read from a single row
// of the Ranking's file.
func updateDb(ctx context.Context, data []string, f Format, gID int) {
	var pID int
	// Add the player information only if he doesn't already exist in the db
	if !db.ContainsPlayer(ctx, data[f.Name], data[f.Surname]) {
		pID = db.Add(ctx, "giocatore", data[f.Name], data[f.Surname])
	} else {
		pID = db.RetrievePlayerID(ctx, data[f.Name], data[f.Surname])
	}
	// Sometime the time is not specified in the ranking file
	time, err := strconv.Atoi(data[f.Time])
	if err != nil || f.Time != -1 {
		time = 0
	}

	rID := db.Add(ctx, "risultato", time, data[f.Exercises], data[f.Score])
	db.Add(ctx, "partecipazione", pID, gID, rID)
}
*/
