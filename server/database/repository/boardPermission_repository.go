package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type BoardPermissionRepository interface {
	Repository[models.BoardPermission]
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
