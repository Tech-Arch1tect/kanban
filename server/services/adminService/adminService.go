package adminService

import (
	"errors"
	"server/database"
	"server/models"
)

type PaginationResult struct {
	Users        []models.User
	TotalRecords int
	TotalPages   int
}

func RemoveUser(userID uint) error {
	return database.DB.UserRepository.Delete(userID)
}

func ListUsers(page, pageSize int, search string) (PaginationResult, error) {
	users, totalRecords, err := database.DB.UserRepository.PaginatedSearch(page, pageSize, search)
	if err != nil {
		return PaginationResult{}, err
	}
	totalPages := (int(totalRecords) + pageSize - 1) / pageSize
	return PaginationResult{
		Users:        users,
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
	}, nil
}

func UpdateUserRole(userID uint, role models.Role) (models.User, error) {
	user, err := database.DB.UserRepository.GetByID(userID)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	if role == "" {
		return models.User{}, errors.New("invalid role")
	}

	user.Role = role
	if err := database.DB.UserRepository.Update(&user); err != nil {
		return models.User{}, errors.New("failed to update user role")
	}

	return user, nil
}
