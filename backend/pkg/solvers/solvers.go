package solvers

import (
	"fmt"
)

// SlimParticipation represents a simplified participation relationship
type SlimParticipation struct {
	City         string  `json:"city"`
	Category     string  `json:"category"`
	IsParis      bool    `json:"isParis"`
	Year         int     `json:"year"`
	MaxExercises int     `json:"eMax"`
	MaxScore     int     `json:"dMax"`
	Score        int     `json:"d"`
	Exercises    int     `json:"e"`
	Time         int     `json:"time"`
	Position     int     `json:"position"`
	PseudoReelo  float64 `json:"pseudoReelo"`
}

// History is a collection of simplified participations
type History []SlimParticipation

// Solver is an array of histories
type Solver History

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
func (ss *Solvers) NewSolver(val SlimParticipation) {
	solver := Solver{val}
	ss.solvers = append(ss.solvers, solver)
	ss.size++
	ss.current = len(ss.solvers) - 1
}

// AppendToCurrent appends the given value to the current solver
func (ss *Solvers) AppendToCurrent(val SlimParticipation) {
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
// The value's consistency is checked against the last inserted value
// in the current solver. We expect to get the results in chronological order.
func (s *Solver) CanAccept(val SlimParticipation) bool {
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
	boundaries, ok := deltaMap[lastResult.Category][val.Category]
	if !ok {
		return false
	}
	if deltaYears > boundaries.Max || deltaYears < boundaries.Min {
		return false
	}

	return true
}

// ShouldBeManual returns whether the Solvers should be manually solved or
// if the software can securely do it on its own
func (ss *Solvers) ShouldBeManual() bool {
	old := ss.curr()
	defer func() { ss.current = old }()

	ss.ResetCursor()
	for ss.Next() {
		curr := *ss.Current()
		for i, s := range curr {
			if i+1 < len(curr)-1 {
				// City changes
				if s.City != curr[i+1].City && s.City != "" && curr[i+1].City != "" {
					return true
				}
				// Distance between years is too big
				if s.Year-curr[i+1].Year > 5 {
					return true
				}
			}

			// I can re-arrange solver in at least two plausible ways
			if ss.curr()+1 < ss.Size() {
				next := ss.solvers[ss.current+1]
				for _, elem := range next {
					if curr.CanAcceptAnyOrder(elem) {
						return true
					}
				}
			}
		}
	}
	return false
}

// ToHistory converts a solver to an history
func (s *Solver) ToHistory() History {
	return History(*s)
}

// CanAcceptAnyOrder checks if a given value belongs to the current solver.
// The value's consistency is checked against every value in the current solver.
func (s *Solver) CanAcceptAnyOrder(val SlimParticipation) bool {
	history := *s
	// If there are no results, accept anything
	if len(history) == 0 {
		return true
	}

	var isAcceptable bool
	for _, result := range history {
		// Do not accept two results in one year
		if result.Year == val.Year && !(result.IsParis || val.IsParis) {
			return false
		}
		// Do not accept two results from paris in the same year
		if result.Year == val.Year && (result.IsParis && val.IsParis) {
			return false
		}

		// Categories growth
		deltaYears := val.Year - result.Year
		boundaries, ok := deltaMap[result.Category][val.Category]
		if !ok {
			continue
		}
		if deltaYears > boundaries.Max || deltaYears < boundaries.Min {
			continue
		}
		isAcceptable = true
	}

	return isAcceptable
}
