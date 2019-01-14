package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const RANK_PATH = "../ranks"

type dataLine struct {
	Name      string
	Surname   string
	City      string
	Exercises int
	Points    int
	Time      int
	Category  string
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

	// Modifying lines in order to work around human errors
	expectedSize = len(format)
	log.Printf("Reading file of year %d, category %s\n", year, category)
	log.Printf("Expected line length is: %d\n", expectedSize)
	log.Printf("Format is: %v\n", format)
	r, err := RunRewriters(Rews, bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	//io.Copy(os.Stdout, r)

	// Parsing each line to save it into the right struct
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		singleLine := parseLine(format, scanner.Text())
		singleLine.Category = category

		results[year] = append(results[year], singleLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLine(format Format, input string) dataLine {
	splitted := strings.Split(input, " ")
	var result dataLine

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
			case "città", "città(provincia)":
				city := strings.Split(strings.Title(value), "(")
				result.City = city[0]
			default:
				log.Println("Unsupported format", fName)
			}
			if err != nil {
				log.Printf("Len of the line is %d", len(splitted))
				log.Fatalf("Could not convert data. The input is: '%s' %v", input, err)
			}
		}
	} else {
		fmt.Println(input)
	}
	return result
}
