package parse

import (
	"fmt"
	"strings"
)

// possibleFormats is just a quick hash-set
var possibleFormats = map[string]struct{}{
	"cognome":          struct{}{},
	"nome":             struct{}{},
	"esercizi":         struct{}{},
	"punti":            struct{}{},
	"tempo":            struct{}{},
	"città":            struct{}{},
	"città(provincia)": struct{}{},
}

// yrFrmt pairs a year with its format
type yrFrmt struct {
	Year int    `json:"year"`
	Frmt string `json:"format"`
}

// allFormats represents an array of yrFrmts
type allFormats struct {
	Formats []yrFrmt `json:"formats"`
}

// Format represents the format used in a ranking file.
// Each element of the map links to the index of the column's number.
type Format map[string]int

// NewFormat returns a new format based on the specified slice of strings
func NewFormat(input []string) (Format, error) {
	format := make(map[string]int)

	for index, value := range input {
		value = strings.ToLower(value)
		if _, ok := possibleFormats[value]; !ok {
			err := fmt.Errorf("format value %s not recognized", value)
			return nil, err
		}
		if value == "città(provincia)" {
			value = "città"
		}
		format[value] = index
	}
	return format, nil
}
