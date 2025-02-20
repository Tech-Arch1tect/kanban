package comment

import (
	"server/database/repository"
	"server/models"
)

func (cs *CommentService) GetComment(commentID uint) (models.Comment, error) {
	comment, err := cs.db.CommentRepository.GetByID(commentID, repository.WithPreload("Task"), repository.WithPreload("User"), repository.WithPreload("Reactions"), repository.WithPreload("Reactions.User"))
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (cs *CommentService) GetCommentReaction(reactionID uint) (models.Reaction, error) {
	reaction, err := cs.db.CommentReactionRepository.GetByID(reactionID, repository.WithPreload("Comment"), repository.WithPreload("User"), repository.WithPreload("Comment.Task"))
	if err != nil {
		return models.Reaction{}, err
	}

	return reaction, nil
}
