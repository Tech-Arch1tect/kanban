package comment

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (cs *CommentService) DeleteCommentRequest(userID, commentID uint) (models.Comment, error) {
	comment, err := cs.db.CommentRepository.GetByID(commentID, repository.WithPreload("Task"))
	if err != nil {
		return models.Comment{}, err
	}

	can, err := cs.rs.CheckRole(userID, comment.Task.BoardID, role.MemberRole)
	if err != nil {
		return models.Comment{}, err
	}

	if !can {
		return models.Comment{}, errors.New("forbidden")
	}

	err = cs.DeleteComment(comment.ID)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (cs *CommentService) DeleteComment(commentID uint) error {
	reactions, err := cs.db.CommentReactionRepository.GetAll(repository.WithWhere("comment_id = ?", commentID))
	if err != nil {
		return err
	}

	for _, reaction := range reactions {
		err = cs.DeleteCommentReaction(reaction.ID)
		if err != nil {
			return err
		}
	}

	return cs.db.CommentRepository.Delete(commentID)
}

func (cs *CommentService) DeleteCommentReactionRequest(userID, reactionID uint) (models.Reaction, error) {
	reaction, err := cs.GetCommentReaction(reactionID)
	if err != nil {
		return models.Reaction{}, err
	}

	can, err := cs.rs.CheckRole(userID, reaction.Comment.Task.BoardID, role.MemberRole)
	if err != nil {
		return models.Reaction{}, err
	}

	if !can {
		return models.Reaction{}, errors.New("forbidden")
	}

	user, err := cs.db.UserRepository.GetByID(userID)
	if err != nil {
		return models.Reaction{}, err
	}

	if user.ID != reaction.UserID && user.Role != models.RoleAdmin {
		return models.Reaction{}, errors.New("forbidden")
	}

	err = cs.DeleteCommentReaction(reaction.ID)
	if err != nil {
		return models.Reaction{}, err
	}

	return reaction, nil
}

func (cs *CommentService) DeleteCommentReaction(reactionID uint) error {
	return cs.db.CommentReactionRepository.HardDelete(reactionID)
}
