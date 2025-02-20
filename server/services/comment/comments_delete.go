package comment

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (cs *CommentService) DeleteComment(userID, commentID uint) (models.Comment, error) {
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

	if err := cs.db.CommentRepository.Delete(comment.ID); err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (cs *CommentService) DeleteCommentReaction(userID, reactionID uint) (models.Reaction, error) {
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

	if err := cs.db.CommentReactionRepository.HardDelete(reaction.ID); err != nil {
		return models.Reaction{}, err
	}

	return reaction, nil
}
