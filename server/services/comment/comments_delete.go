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
