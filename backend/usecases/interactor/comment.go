package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/usecases"
)

// AddComment creates or updates a comment for the given player
func (i *Interactor) AddComment(namesake usecases.Namesake) error {
	ctx := context.Background()
	commentExists := i.CommentRepository.CheckExistenceByPlayerID(ctx, namesake.Comment.Player.ID)
	if commentExists {
		err := i.CommentRepository.UpdateTextByPlayerID(ctx, namesake.Comment.Text, namesake.Comment.Player.ID)
		if err != nil {
			return err
		}
	} else {
		if _, err := i.CommentRepository.Store(ctx, namesake.Comment); err != nil {
			return err
		}
	}
	return nil
}
