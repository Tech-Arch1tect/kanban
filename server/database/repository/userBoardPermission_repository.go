package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type UserBoardPermissionRepository interface {
	Repository[models.UserBoardPermission]
}

type GormUserBoardPermissionRepository struct {
	*GormRepository[models.UserBoardPermission]
}

func NewUserBoardPermissionRepository(db *gorm.DB) UserBoardPermissionRepository {
	return &GormUserBoardPermissionRepository{
		GormRepository: NewGormRepository[models.UserBoardPermission](db),
	}
}

func (r *GormUserBoardPermissionRepository) Migrate() error {
	return r.db.AutoMigrate(&models.UserBoardPermission{})
}
