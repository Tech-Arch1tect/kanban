package permissions

import (
	"server/database"
	"server/models"
)

func CanCreateTask(user models.User, boardID uint) bool {
	return CanAccessBoard(user, boardID)
}

func CanEditTask(user models.User, taskID uint) bool {
	task, err := database.DB.TaskRepository.GetByID(taskID)
	if err != nil {
		return false
	}

	return CanAccessBoard(user, task.BoardID)
}

func CanDeleteTask(user models.User, taskID uint) bool {
	task, err := database.DB.TaskRepository.GetByID(taskID)
	if err != nil {
		return false
	}

	return CanAccessBoard(user, task.BoardID)
}
