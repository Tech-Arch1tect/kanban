package services

import (
	"errors"
	"server/database/repository"
	"server/models"

	"gorm.io/gorm"
)

const (
	MemberRole = "member"
	AdminRole  = "admin"
	ReaderRole = "reader"
)

var roles = []string{MemberRole, AdminRole, ReaderRole}

type RoleService struct {
	db *repository.Database
}

func NewRoleService(db *repository.Database) *RoleService {
	return &RoleService{
		db: db,
	}
}

func (rs *RoleService) SeedRoles() error {
	for _, role := range roles {
		_, err := rs.db.BoardRoleRepository.GetFirst(repository.WithWhere("name = ?", role))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := rs.db.BoardRoleRepository.Create(&models.BoardRole{Name: role}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (rs *RoleService) CheckRole(userID, boardID uint, roleName string) (bool, error) {
	role, err := rs.db.BoardRoleRepository.GetFirst(repository.WithWhere("name = ?", roleName))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	_, err = rs.db.UserBoardRoleRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ? AND board_role_id = ?", userID, boardID, role.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (rs *RoleService) AssignRole(userID, boardID, roleID uint) error {
	existingRole, err := rs.db.UserBoardRoleRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ?", userID, boardID))
	if err == nil {
		if err := rs.db.UserBoardRoleRepository.Delete(existingRole.ID); err != nil {
			return err
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	userBoardRole := models.UserBoardRole{
		UserID:      userID,
		BoardID:     boardID,
		BoardRoleID: roleID,
	}

	return rs.db.UserBoardRoleRepository.Create(&userBoardRole)
}

func (rs *RoleService) RemoveRole(userID, boardID uint) error {
	role, err := rs.db.UserBoardRoleRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ?", userID, boardID))
	if err != nil {
		return err
	}

	return rs.db.UserBoardRoleRepository.Delete(role.ID)
}

func (rs *RoleService) GetRoleByUserAndBoard(userID, boardID uint) (models.BoardRole, error) {
	userBoardRole, err := rs.db.UserBoardRoleRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ?", userID, boardID))
	if err != nil {
		return models.BoardRole{}, err
	}

	role, err := rs.db.BoardRoleRepository.GetFirst(repository.WithWhere("id = ?", userBoardRole.BoardRoleID))
	if err != nil {
		return models.BoardRole{}, err
	}

	return role, nil
}

func (rs *RoleService) GetRolesByBoard(boardID uint) ([]models.BoardRole, error) {
	return rs.db.BoardRoleRepository.GetAll(repository.WithWhere("board_id = ?", boardID))
}
