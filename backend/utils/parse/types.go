package parse

// This is used in magic.go
var expectedSize int

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
}

// DataAll represents a collection of data divided by year
// map[year][]User
type DataAll map[int][]LineInfo
