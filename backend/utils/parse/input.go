package parse

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// readRankingFile reads a ranking from the correct file using the specified
// format. The file's name must be in the format of "year_category.txt"
func readRankingFile(year int, category string, format Format) ([]LineInfo, error) {
	var result []LineInfo
	filePath := fmt.Sprintf("%s/%d/%d_%s.txt", RankPath, year, year, category)
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	result, err = File(bufio.NewReader(file), format, year, category)
	if err != nil {
		return result, err
	}

	return result, nil
}

func parseLine(format Format, input string) (LineInfo, []string) {
	input = strings.ToLower(input)
	splitted := strings.Split(input, " ")
	var result LineInfo

	var index int
	var deltaName int
	var deltaCity int
	var deltaSurname int
	var errs []string

	if input == "" {
		return result, errs
	}

	if len(splitted) < len(format) {
		fmt.Println(input)
		return result, errs
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
					result.Surname = strings.Title(splitted[index])
					for _, c := range commonSurnamePrefix {
						if splitted[index] == c {
							//log.Printf("Line with multi word surname found. Prefix is %s.", c)
							//log.Printf("Line is %v", splitted)

							// Assuming the surname has only 2 words.
							deltaSurname = 1
							value := extractValue(fName, index, deltaSurname, splitted, result)
							result.Surname = strings.Title(value)
						}
					}
					// TODO: not sure about this
					for _, c := range doubleWordSurnames {
						if strings.Contains(input, " "+c+" ") ||
							strings.Contains(input, c+" ") {
							log.Printf("Line with multi word surname found. Surname is %s.", c)

							deltaSurname = len(strings.Split(c, " ")) - 1
							value := extractValue(fName, index, deltaSurname, splitted, result)
							result.Name = strings.Title(value)
						}
					}

				case "nome":
					result.Name = strings.Title(splitted[index])
					for _, c := range doubleWordNames {
						// because of 'DE MARIA LAURA' I inserted the flag, but this excludes whoever has double surname/name
						exceptions := []string{
							"de maria",
							"de marco",
						}
						var excFlag bool
						for _, exc := range exceptions {
							if strings.Contains(input, exc) {
								excFlag = true
							}
						}

						if (strings.Contains(input, " "+c+" ") ||
							strings.Contains(input, c+" ")) &&
							!excFlag {
							log.Printf("Line with multi word name found. Name is %s.", c)
							//log.Printf("Line is %v", splitted)

							deltaName = len(strings.Split(c, " ")) - 1
							value := extractValue(fName, index, deltaName, splitted, result)
							result.Name = strings.Title(value)
						}
					}

				case "esercizi":
					result.Exercises, err = strconv.Atoi(splitted[index])

				case "punti":
					result.Points, err = strconv.Atoi(splitted[index])

				case "tempo":
					result.Time, err = strconv.Atoi(splitted[index])

				case "città", "città(provincia)":
					result.City = strings.Title(splitted[index])
					for _, c := range doubleNameCities {
						if strings.Contains(input, strings.ToLower(c)) {
							//log.Printf("Line with multi word city found. City is %s.", c)
							//log.Printf("Line is %v", splitted)

							if strings.Contains(input, "finalista parigi") {
								r := regexp.MustCompile("finalista parigi [0-9]{4}")
								if r.MatchString(input) {
									deltaCity = 2
								} else {
									deltaCity = 1
								}
							} else {
								deltaCity = len(strings.Split(c, " ")) - 1
							}
							value := extractValue(fName, index, deltaCity, splitted, result)
							result.City = strings.Title(value)
						}
					}

				default:
					log.Println("Unsupported format", fName)
				}
				if err != nil {
					e := fmt.Sprintf("Could not convert data: %v\nThe input is: %v\n", err, input)
					log.Printf(e)
					errs = append(errs, e)
				}
			}
		}
	}
	return result, errs
}

func extractValue(fName string, index, delta int, splitted []string, result LineInfo) string {
	value := splitted[index]
	for i := 1; i < delta+1; i++ {
		value = value + " " + splitted[index+i]
	}

	return value
}
