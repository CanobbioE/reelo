package elo

import "strings"

// Category represents the numerical value for a category
type Category int

// Enum to compare categories
const (
	CE Category = iota
	CM
	C1
	C2
	L1
	L2
	GP
	HC = iota - 1 // H.C. and GP are considered the same
)

var toString = []string{"CE", "CM", "C1", "C2", "L1", "GP", "L2", "HC"}

func (c Category) String() string {
	return toString[c]
}

// CategoryFromString returns a Category from the given string
func CategoryFromString(s string) Category {
	for i, c := range toString {
		if strings.ToUpper(c) == s {
			return Category(i)
		}
	}
	return -1
}
