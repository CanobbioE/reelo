package dto

import "github.com/CanobbioE/reelo/backend/utils/solvers"

// Namesake contains useful details to descern namesakes
type Namesake struct {
	Name     string         `json:"name"`
	Surname  string         `json:"surname"`
	PlayerID int            `json:"playerID"`
	Solver   solvers.Solver `json:"solver"`
	ID       int            `json:"id"`
}
