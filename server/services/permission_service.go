package services

import (
	"errors"
	"server/database/repository"
	"server/models"

	"gorm.io/gorm"
)

const (
	ViewPermission   = "view"
	EditPermission   = "edit"
	DeletePermission = "delete"
)

var permissions = []string{ViewPermission, EditPermission, DeletePermission}

type PermissionService struct {
	db *repository.Database
}

func NewPermissionService(db *repository.Database) *PermissionService {
	return &PermissionService{
		db: db,
	}
}

func (ps *PermissionService) SeedPermissions() error {
	for _, permission := range permissions {
		_, err := ps.db.BoardPermissionRepository.GetFirst(repository.WithWhere("name = ?", permission))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := ps.db.BoardPermissionRepository.Create(&models.BoardPermission{Name: permission}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ps *PermissionService) CheckPermission(userID, boardID uint, permissionName string) (bool, error) {
	permission, err := ps.db.BoardPermissionRepository.GetFirst(repository.WithWhere("name = ?", permissionName))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	_, err = ps.db.UserBoardPermissionRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ? AND permission_id = ?", userID, boardID, permission.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (ps *PermissionService) GrantPermission(userID, boardID, permissionID uint) error {
	userBoardPermission := models.UserBoardPermission{
		UserID:            userID,
		BoardID:           boardID,
		BoardPermissionID: permissionID,
	}

	return ps.db.UserBoardPermissionRepository.Create(&userBoardPermission)
}

func (ps *PermissionService) RevokePermission(userID, boardID, permissionID uint) error {
	userBoardPermission, err := ps.db.UserBoardPermissionRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ? AND permission_id = ?", userID, boardID, permissionID))
	if err != nil {
		return err
	}

	return ps.db.UserBoardPermissionRepository.Delete(userBoardPermission.ID)
}

func (ps *PermissionService) GetPermissionsByUserAndBoard(userID, boardID uint) ([]models.BoardPermission, error) {
	return ps.db.BoardPermissionRepository.GetAll(repository.WithWhere("user_id = ? AND board_id = ?", userID, boardID))
}

func (ps *PermissionService) GetPermissionsByUser(userID uint) ([]models.BoardPermission, error) {
	return ps.db.BoardPermissionRepository.GetAll(repository.WithWhere("user_id = ?", userID))
}

func (ps *PermissionService) GetPermissionsByBoard(boardID uint) ([]models.BoardPermission, error) {
	return ps.db.BoardPermissionRepository.GetAll(repository.WithWhere("board_id = ?", boardID))
}
