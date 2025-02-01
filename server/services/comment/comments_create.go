package comment

import (
	"errors"
	"server/models"
	"server/services/role"
)

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
