package category

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

var toString = []string{"CE", "CM", "C1", "C2", "L1", "L2", "GP", "HC"}

func (c Category) String() string {
	return toString[c]
}

// FromString returns a Category from the given string
func FromString(s string) Category {
	for i, c := range toString {
		if strings.ToUpper(c) == s {
			return Category(i)
		}
	}
	return -1
}

// MaxCategory returns a new category set to the maximum value
func MaxCategory() Category {
	return Category(len(toString) - 1)
}

// MinCategory returns a new category set to the minimum value
func MinCategory() Category {
	return Category(0)
}
