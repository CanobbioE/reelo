package parse

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type DataAll []User

// TODO: really don't like this, but I'm tired. Will move it around later.
var results DataAll

func init() {
	log.Println("Getting cities")
	getCities()
	log.Println("Got cities")
}

func GetUsers() DataAll {
	/*
			TODO: this doesn't work if a category is not present
		 	or if there is another category in the folder.
			We could use readDir for this without enforcing the folder structure

			files, err := ioutil.ReadDir(".")
		    if err != nil {
		        log.Fatal(err)
		    }

		    for _, file := range files {
		        fmt.Println(file.Name())
		    }

			But for now I'm going to just iterate dumbly
	*/

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
	//fmt.Print(results)
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
