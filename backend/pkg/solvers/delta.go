package solvers

import "math"

// Range represents the two ends of integers range
type Range struct {
	Min int
	Max int
}

// DeltaMap represents the inclusive min/max boundries between two categories
/* CATEGORY GROWTH:
ce: 0
cm: 1,2
C1: 3,4
C2: 4,5
L1: 6,7,8
L2: 9, 10, 11
GP: 12,...,99
if two consecutive GP, then MUST do HC
HC: 12,...,99
*/
var deltaMap = map[string]map[string]Range{
	"CE": map[string]Range{
		"CE": Range{0, 0},
		"CM": Range{1, 2},
		"C1": Range{3, 4},
		"C2": Range{5, 6},
		"L1": Range{7, 9},
		"L2": Range{10, 12},
		"GP": Range{13, math.MaxInt32},
		"HC": Range{13, math.MaxInt32},
	},
	"CM": map[string]Range{
		"CM": Range{0, 1},
		"C1": Range{1, 3},
		"C2": Range{3, 5},
		"L1": Range{5, 8},
		"L2": Range{8, 11},
		"GP": Range{11, math.MaxInt32},
		"HC": Range{11, math.MaxInt32},
	},
	"C1": map[string]Range{
		"C1": Range{0, 1},
		"C2": Range{1, 3},
		"L1": Range{3, 6},
		"L2": Range{6, 9},
		"GP": Range{9, math.MaxInt32},
		"HC": Range{9, math.MaxInt32},
	},
	"C2": map[string]Range{
		"C2": Range{0, 1},
		"L1": Range{1, 4},
		"L2": Range{4, 7},
		"GP": Range{7, math.MaxInt32},
		"HC": Range{7, math.MaxInt32},
	},
	"L1": map[string]Range{
		"L1": Range{0, 2},
		"L2": Range{1, 5},
		"GP": Range{4, math.MaxInt32},
		"HC": Range{4, math.MaxInt32},
	},
	"L2": map[string]Range{
		"L2": Range{0, 2},
		"GP": Range{1, math.MaxInt32},
		"HC": Range{1, math.MaxInt32},
	},
	"GP": map[string]Range{
		"GP": Range{0, math.MaxInt32},
		"HC": Range{0, math.MaxInt32},
	},
	"HC": map[string]Range{
		"GP": Range{0, math.MaxInt32},
		"HC": Range{0, math.MaxInt32},
	},
}
