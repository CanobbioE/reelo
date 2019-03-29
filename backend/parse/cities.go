package parse

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var cities []string
var doubleNameCities []string

func getCities() {
	file, err := os.Open("locations")
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
