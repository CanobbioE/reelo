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
	Category  string
}

// parseRankingFile reads a ranking from the correct file using the specified
// format. The file's name must be in the format of "year_category.txt"
func readRankingFile(year int, category string, format Format) dataAll {
	filePath := fmt.Sprintf("%s/%d/%d_%s.txt", RANK_PATH, year, year, category)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Couldn't open file.", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		singleLine := parseLine(format, scanner.Text())
		singleLine.Category = category

		results[year] = append(results[year], singleLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}

func parseLine(format Format, input string) dataLine {
	splitted := strings.Split(input, " ")
	var result dataLine

	//fmt.Println(input)

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
				log.Fatal("Could not convert data. The input is: ", input, err)
			}
		}
		//fmt.Println(result)
	} else {
		fmt.Println(input)
		//log.Print("Not implemented")
	}
	//fmt.Println()
	return result
}
