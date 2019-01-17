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

var expectedSize int

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
	//input = strings.ToLower(input)
	splitted := strings.Split(input, " ")
	var result dataLine

	// Handling base case.
	// Otherwise it means there are fields with two or more words.
	if len(splitted) == len(format) {
		for fName, index := range format {
			result = assignField(splitted[index], fName, result)
		}
	} else {
		isOnlyDoubleCity := false
		isOnlyDoubleName := false

		for _, c := range doubleNameCities {
			if strings.Contains(input, c) {
				cIndex, ok := format["città"]
				cWords := len(strings.Split(c, " "))

				// checking if the multi word city is the only parameter with more than one word
				log.Printf("Line with multi word city found. City is %s.", c)

				if len(splitted) == len(format)+cWords-1 {
					log.Printf("Should be only parameter w/ multiple words.")
					log.Printf("Line is %v", splitted)
					isOnlyDoubleCity = true

					for fName, fIndex := range format {
						// Handling index that come after the double words
						index := fIndex
						if fIndex > cIndex && ok {
							index = fIndex + cWords - 1
						}

						value := splitted[index]
						if fName == "città" {
							for i := 1; i < cWords; i++ {
								value = value + " " + splitted[index+i]
							}
						}
						result = assignField(value, fName, result)
					}
				}
			}
		}

		for _, n := range doubleWordNames {
			if strings.Contains(strings.ToLower(input), strings.ToLower(n)) {
				nIndex, ok := format["nome"]
				nWords := len(strings.Split(n, " "))

				// checking if the multi word name is the only parameter with more than one word
				log.Printf("Line with multi word name found. Name is %s.", n)

				if len(splitted) == len(format)+nWords-1 {
					log.Printf("Should be only parameter w/ multiple words.")
					log.Printf("Line is %v", splitted)
					isOnlyDoubleName = true

					for fName, fIndex := range format {
						// Handling index that come after the double words
						index := fIndex
						if fIndex > nIndex && ok {
							index = fIndex + nWords - 1
						}

						value := splitted[index]
						if fName == "nome" {
							for i := 1; i < nWords; i++ {
								value = value + " " + splitted[index+i]
							}
						}
						result = assignField(value, fName, result)
					}
				}
			}
		}

		if !isOnlyDoubleCity && !isOnlyDoubleName {
			log.Printf("Found an exception. Len got: %d, exp %d.", len(splitted), len(format))
			fmt.Println(input)
		}
	}
	return result
}

func assignField(value, fName string, result dataLine) dataLine {
	var err error
	switch fName {
	case "cognome":
		result.Surname = strings.Title(strings.ToLower(value))
	case "nome":
		result.Name = strings.Title(strings.ToLower(value))
	case "esercizi":
		result.Exercises, err = strconv.Atoi(value)
	case "punti":
		result.Points, err = strconv.Atoi(value)
	case "tempo":
		result.Time, err = strconv.Atoi(value)
	case "città", "città(provincia)":
		result.City = strings.Title(strings.ToLower(value))
	default:
		log.Println("Unsupported format", fName)
	}
	if err != nil {
		log.Printf("Could not convert data. The input is: '%s' %v", value, err)
	}
	return result
}
