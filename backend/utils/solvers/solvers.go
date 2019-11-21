package solvers

import (
	"fmt"

	rdb "github.com/CanobbioE/reelo/backend/db"
)

// Solver is an array of histories
type Solver []rdb.History

// Solvers is a an array of Solvers, which are arrays of histories.
// The solvers length is equal to the number of plausible histories found.
type Solvers struct {
	solvers []Solver
	size    int
	current int
}

// New creates a new Solvers with one empty solver
// Due to the usage we want to do with this type, the current position is set to -1
func New() *Solvers {
	return &Solvers{
		solvers: make([]Solver, 1, 1),
		size:    1,
		current: -1,
	}
}

// String stringify the solvers
func (ss *Solvers) String() string {
	var ret string
	for _, s := range ss.solvers {
		ret += fmt.Sprintf("%v\n", s)
	}
	return ret
}

func (ss *Solvers) curr() int {
	if ss.current < 0 {
		return 0
	}
	return ss.current
}

// Next advance the current position and returns true if there's a next
// solver, false otherwise. The cycle start from position 0 so we can directly
// call for solvers.Next() without saving the current solver before iterating.
func (ss *Solvers) Next() bool {
	if ss.current+1 >= ss.size {
		return false
	}
	ss.current = ss.current + 1
	return ss.current < ss.size
}

// Current returns the pointer to the current solver
func (ss *Solvers) Current() *Solver {
	return &ss.solvers[ss.curr()]
}

// ResetCursor resets the cursor to the initial position
func (ss *Solvers) ResetCursor() {
	ss.current = -1
}

// Size returns the number of solvers in the underlying array
func (ss *Solvers) Size() int {
	return ss.size
}

// SetCurrent assigns a solver to the current cursor position
func (ss *Solvers) SetCurrent(s Solver) {
	ss.solvers[ss.curr()] = s
}

// NewSolver allocates a new solver containing one value and moves the cursor
// to the newly added value
func (ss *Solvers) NewSolver(val rdb.History) {
	solver := []rdb.History{val}
	ss.solvers = append(ss.solvers, solver)
	ss.size++
	ss.current = len(ss.solvers) - 1
}

// AppendToCurrent appends the given value to the current solver
func (ss *Solvers) AppendToCurrent(val rdb.History) {
	ss.solvers[ss.curr()] = append(ss.solvers[ss.curr()], val)
}

// HasNext returns true if there's a next solver in the array, false otherwise
func (ss *Solvers) HasNext() bool {
	if ss.current+1 >= ss.size {
		return false
	}
	return true
}

// CanAccept checks if a given value belongs to the current solver.
// The value's cosistency is checked against the last inserted value
// in the current solver. We expect to get the results in chronological order.
func (s *Solver) CanAccept(val rdb.History) bool {
	history := *s
	// If there are no results, accept anything
	if len(history) == 0 {
		return true
	}

	for _, result := range history {
		// Do not accept two results in one year
		if result.Year == val.Year && !(result.IsParis || val.IsParis) {
			return false
		}
		// Do not accept two results from paris in the same year
		if result.Year == val.Year && (result.IsParis && val.IsParis) {
			return false
		}
	}

	// Categories growth
	lastResult := history[len(history)-1]
	deltaYears := val.Year - lastResult.Year
	boundries, ok := deltaMap[lastResult.Category][val.Category]
	if !ok {
		return false
	}
	if deltaYears > boundries.Max || deltaYears < boundries.Min {
		return false
	}

	// TODO: check places

	return true
}
