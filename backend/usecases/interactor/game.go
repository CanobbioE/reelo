package interactor

import "context"

// ListYears returns an array with all the years that are stored in the repository
func (i *Interactor) ListYears() ([]int, error) {
	return i.GameRepository.FindAllYears(context.Background())
}
