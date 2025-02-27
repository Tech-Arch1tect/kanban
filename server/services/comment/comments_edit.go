package comment

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (cs *CommentService) EditComment(user *models.User, commentID uint, text string) (models.Comment, models.Comment, error) {
	comment, err := cs.db.CommentRepository.GetByID(commentID, repository.WithPreload("Task"))
	if err != nil {
		return models.Comment{}, models.Comment{}, err
	}

	if user.Role != models.RoleAdmin && user.ID != comment.UserID {
		return models.Comment{}, models.Comment{}, errors.New("forbidden")
	}

	can, _ := cs.rs.CheckRole(user.ID, comment.Task.BoardID, role.MemberRole)
	if !can {
		return models.Comment{}, models.Comment{}, errors.New("forbidden")
	}

	oldComment := comment
	comment.Text = text
	if err := cs.db.CommentRepository.Update(&comment); err != nil {
		return models.Comment{}, models.Comment{}, err
	}

	return comment, oldComment, nil
}
