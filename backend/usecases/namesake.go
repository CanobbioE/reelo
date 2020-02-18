package usecases

import (
	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils/solvers"
)

// Namesake represents a namesake
type Namesake struct {
	Comment domain.Comment `json:"comment"`
	Solver  solvers.Solver `json:"solver"`
	ID      int            `json:"id"`
	Player  domain.Player  `json:"player"`
}
