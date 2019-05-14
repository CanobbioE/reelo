package parse

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var cities []string
var doubleNameCities []string

// GetCities reads the list of cities from the "locations" file
// and saves it in two arrays based on the number of words it
// is composed by.
// Should be called only once on initialization of the app
func GetCities() {
	file, err := os.Open("./utils/parse/locations")
	if err != nil {
		log.Fatal("Couldn't open file.", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		city := strings.Split(scanner.Text(), ",")

		cities = append(cities, strings.ToLower(city[0]))

		if strings.ContainsAny(city[0], " ") {
			doubleNameCities = append(doubleNameCities, city[0])
		}
	}
}
