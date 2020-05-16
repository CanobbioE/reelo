package usecases

import (
	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/pkg/solvers"
)

// Namesake represents a namesake
// TODO: soon obsolete
type Namesake struct {
	Comment domain.Comment `json:"comment"`
	Solver  solvers.Solver `json:"solver"`
	ID      int            `json:"id"`
	Player  domain.Player  `json:"player"`
}
