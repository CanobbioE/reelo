package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var possibleFormats = map[string]struct{}{
	"cognome":          struct{}{},
	"nome":             struct{}{},
	"esercizi":         struct{}{},
	"punti":            struct{}{},
	"tempo":            struct{}{},
	"città":            struct{}{},
	"città(provincia)": struct{}{},
}

type yrFrmt struct {
	Year int    `json:"year"`
	Frmt string `json:"format"`
}
type allFormats struct {
	Formats []yrFrmt `json:"formats"`
}

// Format represents the format used in a ranking file.
// Each element of the map links to the index of the column's number.
type Format map[string]int

// newFormat returns a new format based on the slice of string passed
func newFormat(input []string) Format {
	format := make(map[string]int)

	for index, value := range input {
		value = strings.ToLower(value)
		if _, ok := possibleFormats[value]; !ok {
			log.Fatalf("format value %s not recognized.", value)
		}
		if value == "città(provincia)" {
			value = "città"
		}
		format[value] = index
	}
	return format
}

func readFormats() allFormats {
	file, err := os.Open(RANK_PATH + "/formats.json")
	if err != nil {
		log.Fatal("Couldn't open formats file.", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var result allFormats

	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Fatal("Couldn't unmarshal formats json.", err)
	}

	return result
}

func retrieveFormat(year int, input allFormats) []string {
	for _, value := range input.Formats {
		if value.Year == year {
			return strings.Split(value.Frmt, ", ")
		}
	}
	return nil
}
