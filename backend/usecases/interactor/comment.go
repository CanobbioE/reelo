package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/usecases"
)

// AddComment creates or updates a comment for the given player
func (i *Interactor) AddComment(namesake usecases.Namesake) error {
	ctx := context.Background()
	if i.CommentRepository.CheckExistenceByPlayerID(ctx, namesake.PlayerID) {
		if err := i.CommentRepository.UpdateTextByPlayerID(ctx, n.Comment, n.PlayerID); err != nil {
			return err
		}
	} else {
		if _, err := i.CommentRepository.Store(ctx, "commenti", namesake.PlayerID, namesake.Comment); err != nil {
			return err
		}
	}
	return nil
}
