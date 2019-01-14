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
				result.City = value
			default:
				log.Println("Unsupported format", fName)
			}
			if err != nil {
				log.Printf("Len of the line is %d", len(splitted))
				log.Printf("Could not convert data. The input is: '%s' %v", input, err)
			}
		}
	} else {
		isDoubleCity := false

		for _, c := range doubleNameCities {
			if strings.Contains(input, c) {
				cIndex := format["città"]
				cWords := len(strings.Split(c, " "))

				// checking if the multi word city is the only parameter with more than one word
				log.Printf("Line with multi word city found. City is %s.", c)

				if len(splitted) == len(format)+cWords-1 {
					log.Printf("Should be only parameter w/ multiple words.")
					log.Printf("Line is %v", splitted)
					isDoubleCity = true

					for fName, fIndex := range format {
						// Handling index that come after the double words
						index := fIndex
						if fIndex > cIndex {
							index = fIndex + cWords - 1
						}

						var err error
						var value string

						switch fName {
						case "cognome":
							value = strings.Title(strings.ToLower(splitted[index]))
							log.Printf("Surname value: %s, index: %d", value, index)
							result.Surname = value
						case "nome":
							value = strings.Title(strings.ToLower(splitted[index]))
							log.Printf("Name value: %s, index: %d", value, index)
							result.Name = value
						case "esercizi":
							result.Exercises, err = strconv.Atoi(splitted[index])
						case "punti":
							result.Points, err = strconv.Atoi(splitted[index])
						case "tempo":
							result.Time, err = strconv.Atoi(splitted[index])
						case "città", "città(provincia)":
							for i := 0; i < cWords; i++ {
								value = value + " " + splitted[index+i]
							}
							value = strings.Title(strings.ToLower(value))
							log.Printf("City value: %s, index: %d", value, index)
							result.City = value
						default:
							log.Println("Unsupported format", fName)
						}
						if err != nil {
							log.Printf("Len of the line is %d", len(splitted))
							log.Printf("Could not convert data. The input is: '%s' %v", input, err)
						}
					}
				}
			}
		}

		if !isDoubleCity {
			log.Printf("Found an exception. Len got: %d, exp %d.", len(splitted), len(format))
			fmt.Println(input)
		}
	}
	return result
}
