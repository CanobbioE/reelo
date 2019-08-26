package elo

import "strings"

type category int

// Enum to compare categories
const (
	CE category = iota
	CM
	C1
	C2
	L1
	GP
	L2
	HC
)

var toString = []string{"CE", "CM", "C1", "C2", "L1", "GP", "L2", "HC"}

func (c category) String() string {
	return toString[c]
}

// CategoryFromString returns a Category from the given string
func categoryFromString(s string) category {
	for i, c := range toString {
		if strings.ToUpper(c) == s {
			return category(i)
		}
	}
	return -1
}
