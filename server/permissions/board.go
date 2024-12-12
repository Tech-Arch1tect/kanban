package permissions

import (
	"server/database"
	"server/models"
)

func CanEditBoard(user models.User, id uint) bool {
	board, err := database.DB.BoardRepository.GetByID(id)
	if err != nil {
		return false
	}

	permission, err := database.DB.BoardRepository.GetPermission(user.ID, board.ID)
	if err != nil {
		return false
	}

	return permission.Edit
}

func CanCreateBoard(user models.User, _ uint) bool {
	return user.Role == models.RoleAdmin
}

func CanDeleteBoard(user models.User, id uint) bool {
	board, err := database.DB.BoardRepository.GetByID(id)
	if err != nil {
		return false
	}

	if board.OwnerID == user.ID {
		return true
	}

	permission, err := database.DB.BoardRepository.GetPermission(user.ID, board.ID)
	if err != nil {
		return false
	}

	return permission.Delete
}

func CanAccessBoard(user models.User, id uint) bool {
	board, err := database.DB.BoardRepository.GetByID(id)
	if err != nil {
		return false
	}

	if board.OwnerID == user.ID {
		return true
	}

	permission, err := database.DB.BoardRepository.GetPermission(user.ID, board.ID)
	if err != nil {
		return false
	}

	if permission != (models.BoardPermission{}) {
		return true
	}

	return false
}
