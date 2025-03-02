package admin

import (
	"errors"
	"server/database/repository"
	"server/models"
)

type AdminService struct {
	db *repository.Database
}

func NewAdminService(db *repository.Database) *AdminService {
	return &AdminService{db: db}
}

type AdminPaginationResult struct {
	Users        []models.User
	TotalRecords int
	TotalPages   int
}

func (s *AdminService) RemoveUser(userID uint) error {
	return s.db.UserRepository.Delete(userID)
}

func (s *AdminService) ListUsers(page, pageSize int, search string) (AdminPaginationResult, error) {
	users, totalRecords, err := s.db.UserRepository.PaginatedSearch(page, pageSize, search, "email", "created_at DESC")
	if err != nil {
		return AdminPaginationResult{}, err
	}
	totalPages := (int(totalRecords) + pageSize - 1) / pageSize
	return AdminPaginationResult{
		Users:        users,
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
	}, nil
}

func (s *AdminService) UpdateUserRole(userID uint, role models.Role) (models.User, error) {
	user, err := s.db.UserRepository.GetByID(userID)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	if role == "" {
		return models.User{}, errors.New("invalid role")
	}

	user.Role = role
	if err := s.db.UserRepository.Update(&user); err != nil {
		return models.User{}, errors.New("failed to update user role")
	}

	return user, nil
}
