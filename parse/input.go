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

	var index int
	var deltaName int
	var deltaCity int
	var deltaSurname int

	if input == "" {
		return result
	}

	if len(splitted) < len(format) {
		fmt.Println(input)
		return result
	}

	// going in numerical order
	// TODO: consider if it's worth inverting format's key value order
	for i := 0; i < len(format); i++ {
		for fName, fIndex := range format {
			if i == fIndex {
				index = fIndex + deltaName + deltaCity + deltaSurname

				var err error
				switch fName {
				case "cognome":
					for _, c := range commonSurnamePrefix {
						if strings.ToLower(splitted[index]) == strings.ToLower(c) {
							log.Printf("Line with multi word surname found. Prefix is %s.", c)
							log.Printf("Line is %v", splitted)

							// Assuming the surname has only 2 words.
							deltaSurname = 1
							value := extractValue(fName, index, deltaSurname, splitted, result)
							result.Surname = strings.Title(strings.ToLower(value))
						}
					}

				case "nome":
					for _, c := range doubleWordNames {
						if strings.Contains(strings.ToLower(input), " "+strings.ToLower(c)+" ") {
							log.Printf("Line with multi word name found. Name is %s.", c)
							log.Printf("Line is %v", splitted)

							deltaName = len(strings.Split(c, " ")) - 1
							value := extractValue(fName, index, deltaName, splitted, result)
							result.Name = strings.Title(strings.ToLower(value))
						}
					}

				case "esercizi":
					result.Exercises, err = strconv.Atoi(splitted[index])

				case "punti":
					result.Points, err = strconv.Atoi(splitted[index])

				case "tempo":
					result.Time, err = strconv.Atoi(splitted[index])

				case "città", "città(provincia)":
					for _, c := range doubleNameCities {
						if strings.Contains(input, c) {
							log.Printf("Line with multi word city found. City is %s.", c)
							log.Printf("Line is %v", splitted)

							deltaCity = len(strings.Split(c, " ")) - 1
							value := extractValue(fName, index, deltaCity, splitted, result)
							result.City = strings.Title(strings.ToLower(value))
						}
					}

				default:
					log.Println("Unsupported format", fName)
				}
				if err != nil {
					log.Printf("Could not convert data. The input is: %v", err)
					fmt.Println(input)
				}
			}
		}
	}
	return result
}

func extractValue(fName string, index, delta int, splitted []string, result dataLine) string {
	value := splitted[index]
	for i := 1; i < delta+1; i++ {
		value = value + " " + splitted[index+i]
	}

	return value
}
