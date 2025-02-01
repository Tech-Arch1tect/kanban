package role

import (
	"errors"
	"server/database/repository"
	"server/models"

	"gorm.io/gorm"
)

func (rs *RoleService) AssignRole(userID, boardID uint, role AppRole) error {
	existingRole, err := rs.db.UserBoardRoleRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ?", userID, boardID))
	if err == nil {
		if err := rs.db.UserBoardRoleRepository.HardDelete(existingRole.ID); err != nil {
			return err
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	roleDb, err := rs.db.BoardRoleRepository.GetFirst(repository.WithWhere("name = ?", role.Name))
	if err != nil {
		return err
	}

	userBoardRole := models.UserBoardRole{
		UserID:      userID,
		BoardID:     boardID,
		BoardRoleID: roleDb.ID,
	}

	return rs.db.UserBoardRoleRepository.Create(&userBoardRole)
}

func (rs *RoleService) RemoveRole(userID, boardID uint) error {
	userBoardRole, err := rs.db.UserBoardRoleRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ?", userID, boardID))
	if err != nil {
		return err
	}

	return rs.db.UserBoardRoleRepository.Delete(userBoardRole.ID)
}
