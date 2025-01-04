package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type UserBoardRoleRepository interface {
	Repository[models.UserBoardRole]
}

type GormUserBoardRoleRepository struct {
	*GormRepository[models.UserBoardRole]
}

func NewUserBoardRoleRepository(db *gorm.DB) UserBoardRoleRepository {
	return &GormUserBoardRoleRepository{
		GormRepository: NewGormRepository[models.UserBoardRole](db),
	}
}

func (r *GormUserBoardRoleRepository) Migrate() error {
	return r.db.AutoMigrate(&models.UserBoardRole{})
}
