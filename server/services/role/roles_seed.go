package role

import (
	"errors"
	"server/database/repository"
	"server/models"

	"gorm.io/gorm"
)

func (rs *RoleService) SeedRoles() error {
	for _, role := range roleMap {
		_, err := rs.db.BoardRoleRepository.GetFirst(repository.WithWhere("name = ?", role.Name))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := rs.db.BoardRoleRepository.Create(&models.BoardRole{Name: role.Name}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
