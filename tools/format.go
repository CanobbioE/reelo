package tools

// Format represents the format used in a ranking file.
// Each field represents an information's index inside a line of the file.
type Format struct {
	Exercises, Score, Time int
	Name, Surname          int
	City                   int
}

// newFormat returns a new format based on the slice of string passed
func newFormat(ff []string) *Format {
	return &Format{
		Name:      indexOf(ff, "nome"),
		Surname:   indexOf(ff, "cognome"),
		City:      indexOf(ff, "sede"),
		Exercises: indexOf(ff, "esercizi"),
		Score:     indexOf(ff, "punteggio"),
		Time:      indexOf(ff, "tempo"),
	}
}

// indexOf returns the position of the pattern's first occurency inside
// the ss slice
func indexOf(ss []string, pattern string) int {
	for i, s := range ss {
		if s == pattern {
			return i
		}
	}
	return -1
}
