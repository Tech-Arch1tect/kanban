package permissions

import (
	"server/database"
	"server/models"
)

func CanCreateComment(user models.User, taskID uint) bool {
	task, err := database.DB.TaskRepository.GetByID(taskID)
	if err != nil {
		return false
	}

	return CanAccessBoard(user, task.BoardID)
}

func CanDeleteComment(user models.User, commentID uint) bool {
	comment, err := database.DB.CommentRepository.GetByID(commentID)
	if err != nil {
		return false
	}

	return comment.UserID == user.ID
}

func CanEditComment(user models.User, commentID uint) bool {
	comment, err := database.DB.CommentRepository.GetByID(commentID)
	if err != nil {
		return false
	}

	return comment.UserID == user.ID
}
