package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/utils"
)

// ListYears returns an array of integers representing all the years that are
// stored in the repository.
func (i *Interactor) ListYears() ([]int, utils.Error) {
	years, err := i.GameRepository.FindAllYears(context.Background())
	if err != nil {
		i.Log("ListYears: cannot find all: %v", err)
		return years, utils.NewError(err, "E_DB_FIND", 500)
	}
	return years, utils.NewNilError()
}
