package inputs

import (
	"fmt"
	"strings"
)

type Results struct {
	Category  string
	Exercises int
	Points    int
	Time      int
}

type Relevant struct {
	MaxPoints, MaxExercises, MaxTime int
	MinPoints, MinExercises, MinTime int
	AvgPoints, AvgExercises, AvgTime int
}

type PersonInfo map[string]map[int]Results

type YearInfo map[int]map[string]Relevant

func (pi PersonInfo) addOneYearInfo(name, surname string, year int, res Results) {
	fullName := strings.ToUpper(fmt.Sprintf("%s %s", name, surname))
	pi[fullName] = map[int]Results{year: res}
}

func (yi YearInfo) addOneCatInfo(year int, cat string, rel Relevant) {
	yi[year] = map[string]Relevant{cat: rel}
}
