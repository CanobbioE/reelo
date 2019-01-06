package main

import (
	"log"
	"strings"
)

var possibleFormats = map[string]interface{}{
	"cognome":          struct{}{},
	"nome":             struct{}{},
	"esercizi":         struct{}{},
	"punti":            struct{}{},
	"tempo":            struct{}{},
	"città":            struct{}{},
	"città(provincia)": struct{}{},
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

		format[value] = index
	}

	return format
}
