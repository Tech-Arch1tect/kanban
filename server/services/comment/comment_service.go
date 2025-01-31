package comment

import (
	"errors"
	"server/database/repository"
	"server/internal/helpers"
	"server/models"
	"server/services/role"
)

type CommentService struct {
	db *repository.Database
	rs *role.RoleService
	hs *helpers.HelperService
}

func NewCommentService(db *repository.Database, rs *role.RoleService, hs *helpers.HelperService) *CommentService {
	return &CommentService{db: db, rs: rs, hs: hs}
}

func (cs *CommentService) CreateComment(userID, taskID uint, text string) (models.Comment, error) {
	task, err := cs.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Comment{}, err
	}

	can, err := cs.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if err != nil {
		return models.Comment{}, err
	}

	if !can {
		return models.Comment{}, errors.New("forbidden")
	}

	comment := models.Comment{
		TaskID: taskID,
		Text:   text,
		UserID: userID,
	}

	if err := cs.db.CommentRepository.Create(&comment); err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

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

func (cs *CommentService) EditComment(user *models.User, commentID uint, text string) (models.Comment, error) {
	comment, err := cs.db.CommentRepository.GetByID(commentID, repository.WithPreload("Task"))
	if err != nil {
		return models.Comment{}, err
	}

	if user.Role != models.RoleAdmin && user.ID != comment.UserID {
		return models.Comment{}, errors.New("forbidden")
	}

	can, _ := cs.rs.CheckRole(user.ID, comment.Task.BoardID, role.MemberRole)
	if !can {
		return models.Comment{}, errors.New("forbidden")
	}

	comment.Text = text
	if err := cs.db.CommentRepository.Update(&comment); err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}
