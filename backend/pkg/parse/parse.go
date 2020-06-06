package parse

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// All parses all files in the ranks folder
func All() (DataAll, error) {
	var results = make(DataAll)
	formats := readFormats()
	years := findYears()
	categories := []string{"C1", "C2", "GP", "L1", "L2"}

	for _, year := range years {
		inputFormat := retrieveFormat(year, formats)
		format, err := NewFormat(inputFormat)
		if err != nil {
			return nil, err
		}
		for _, category := range categories {
			var err error
			// TODO: start = starts[year][cat]
			// TODO: end = ends[year][cat]
			results[year], err = readRankingFile(year, category, format)
			if err != nil {
				log.Printf("Error parsing all files: %v", err)
				return nil, err
			}
			// TODO. results[year].Start = start
			// TODO: results[year].End = End
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
	err := filepath.Walk(RankPath, func(path string, info os.FileInfo, err error) error {
		// TODO: improve this function.
		if path != "./ranks" && path != "./ranks/formats" && len(path) < 14 {
			log.Println(path)
			year, err := strconv.Atoi(path[len("ranks/"):])

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

func prettyPrintErrors(dst string, errs []string) string {
	for _, e := range errs {
		dst = fmt.Sprintf("%v\n%v", dst, e)
	}
	return dst
}
