package parse

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

// RankPath specifies where the ranking files are stored
const RankPath = "./ranks"

// LineInfo represents information contained in a single line of a ranking file
type LineInfo struct {
	Name      string
	Surname   string
	City      string
	Exercises int
	Points    int
	Time      int
	Category  string
	Year      int
	Position  int
	Start     int
	End       int
}

// DataAll represents a collection of data divided by year
// map[year][]Player
type DataAll map[int][]LineInfo

// All parses all files in the ranks folder
func All() (DataAll, error) {
	var results = make(DataAll)
	years := findYears()
	categories := []string{"C1", "C2", "GP", "L1", "L2"}

	for _, year := range years {
		inputFormat := strings.Split(formats[year], " ")
		format, err := NewFormat(inputFormat)
		if err != nil {
			return nil, err
		}
		for _, category := range categories {
			log.Printf("Parsing %d %s...\n", year, category)
			var err error
			results[year], err = readRankingFile(year, category, format)
			if err != nil {
				log.Printf("Error parsing all files: %v", err)
				return nil, err
			}
		}
	}

	return results, nil
}

// File parses the specified file expecting it to be in the given format
func File(fileReader io.Reader, format Format, year int, category string) ([]LineInfo, error) {
	var results []LineInfo
	warning := false
	mergedErrs := fmt.Sprintf("Multiple parsing errors:")
	expectedSize = len(format)

	start := scoreHelp[year][category].Start
	end := scoreHelp[year][category].End

	r, err := RunRewriters(Rews, fileReader)
	if err != nil {
		return results, err
	}

	// Parsing each line to save it into the right struct
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		singleLine, errs := parseLine(format, scanner.Text())
		singleLine.Position = i
		if len(errs) > 0 {
			// Stacking up errors so the user can fix them all in one go
			mergedErrs = prettyPrintErrors(mergedErrs, errs)
			warning = true
		}
		singleLine.Category = category
		singleLine.Year = year
		singleLine.Start = start
		singleLine.End = end

		results = append(results, singleLine)
	}
	if err := scanner.Err(); err != nil {
		return results, err
	}
	if warning {
		return results, fmt.Errorf("%v", mergedErrs)
	}

	return results, nil
}

func findYears() []int {
	var years []int
	for year := 2003; year < 2011; year++ {
		years = append(years, year)
	}
	return years
}

func prettyPrintErrors(dst string, errs []string) string {
	for _, e := range errs {
		dst = fmt.Sprintf("%v\n%v", dst, e)
	}
	return dst
}
