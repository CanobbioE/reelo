package elo

type category int

// Enum to compare categories
const (
	C1 category = iota
	C2
	CE
	CM
	L1
	L2
	GP
	HC
)

var toString = []string{"C1", "C2", "CE", "CM", "L1", "L2", "GP", "HC"}

func (c category) String() string {
	return toString[c]
}

// CategoryFromString returns a Category from the given string
func categoryFromString(s string) category {
	for i, c := range toString {
		if c == s {
			return category(i)
		}
	}
	return -1
}
