package role

import (
	"errors"
	"server/models"

	"gorm.io/gorm"
)

func (rs *RoleService) CheckRole(userID, boardID uint, requiredRole AppRole) (bool, error) {
	user, err := rs.db.UserRepository.GetByID(userID)
	if err != nil {
		return false, err
	}

	if user.Role == models.RoleAdmin {
		return true, nil
	}

	userBoardRole, err := rs.GetRoleByUserAndBoard(userID, boardID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	userAppRole, exists := roleMap[userBoardRole.Name]
	if !exists {
		return false, errors.New("user role not found")
	}

	if userAppRole.AccessLevel >= requiredRole.AccessLevel {
		return true, nil
	}
	return false, nil
}
