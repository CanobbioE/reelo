package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// map[year][]dataLine
type dataAll map[int][]dataLine

// TODO: really don't like this, but I'm tired. Will move it around later.
var results = make(dataAll)

func init() {
	log.Println("Getting cities")
	getCities()
	log.Println("Got cities")
}

func main() {
	formats := readFormats() // []struct{int, string}
	years := findYears()     // []int
	categories := []string{"C1", "C2", "GP", "L1", "L2"}

	for _, year := range years {
		inputFormat := retrieveFormat(year, formats)
		format := newFormat(inputFormat)
		for _, category := range categories {
			readRankingFile(year, category, format)
		}
	}
	// TODO: use result to populate db
}

// findYears walk through the ranking folders and returns an array of integer
// representing all the years that has to be processed.
func findYears() []int {
	var years []int
	err := filepath.Walk(RANK_PATH, func(path string, info os.FileInfo, err error) error {
		// TODO: improve this function.
		if path != "../ranks" && path != "../ranks/formats" && len(path) < 14 {
			year, err := strconv.Atoi(path[9:13])
			if err != nil {
				log.Fatal("Error parsing years.", err)
			}
			years = append(years, year)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return years
}
