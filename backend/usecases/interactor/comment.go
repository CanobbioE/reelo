package interactor

import (
	"context"

	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// AddComment creates or updates a comment for the given player
func (i *Interactor) AddComment(namesake usecases.Namesake) utils.Error {
	ctx := context.Background()
	commentExists := i.CommentRepository.CheckExistenceByPlayerID(ctx, namesake.Comment.Player.ID)
	if commentExists {
		err := i.CommentRepository.UpdateTextByPlayerID(ctx, namesake.Comment.Text, namesake.Comment.Player.ID)
		if err != nil {
			i.Logger.Log("AddComment: cannot update text: %v", err)
			return utils.NewError(err, "E_DB_UPDATE", 500)
		}
	} else {
		if _, err := i.CommentRepository.Store(ctx, namesake.Comment); err != nil {
			i.Logger.Log("AddComment: cannot store comment: %v", err)
			return utils.NewError(err, "E_DB_CREATE", 500)
		}
	}
	return utils.NewNilError()
}
