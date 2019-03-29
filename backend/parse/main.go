package parse

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// DataAll represents a collection of data divided by year
// map[year][]dataLine
type DataAll map[int][]dataLine

// TODO: really don't like this, but I'm tired. Will move it around later.
var results = make(DataAll)

func init() {
	log.Println("Getting cities")
	getCities()
	log.Println("Got cities")
}

// All parses all files in the ranks folder
func All() DataAll {
	formats := readFormats()
	years := findYears()
	categories := []string{"C1", "C2", "GP", "L1", "L2"}

	for _, year := range years {
		inputFormat := retrieveFormat(year, formats)
		format := newFormat(inputFormat)
		for _, category := range categories {
			readRankingFile(year, category, format)
		}
	}

	return results
}

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
