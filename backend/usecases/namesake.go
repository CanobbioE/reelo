package usecases

import (
	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/utils/solvers"
)

type Namesake struct {
	Comment domain.Comment `json:"comment"`
	Solver  solvers.Solver `json:"solver"`
	ID      int            `json:"id"`
}
