package database

import (
	"server/models"

	"gorm.io/gorm"
)

type BoardPermissionRepository interface {
	Repository[models.BoardPermission]
	GetPermissionsByUserID(userID uint) ([]models.BoardPermission, error)
}

type GormBoardPermissionRepository struct {
	*GormRepository[models.BoardPermission]
}

func NewBoardPermissionRepository(db *gorm.DB) BoardPermissionRepository {
	return &GormBoardPermissionRepository{
		GormRepository: NewGormRepository[models.BoardPermission](db),
	}
}

func (r *GormBoardPermissionRepository) Migrate() error {
	return r.db.AutoMigrate(&models.BoardPermission{})
}

func (r *GormBoardPermissionRepository) GetPermissionsByUserID(userID uint) ([]models.BoardPermission, error) {
	var permissions []models.BoardPermission
	result := r.db.Where("user_id = ?", userID).Find(&permissions)
	return permissions, result.Error
}
