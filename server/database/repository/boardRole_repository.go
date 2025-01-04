package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type BoardRoleRepository interface {
	Repository[models.BoardRole]
}

type GormBoardRoleRepository struct {
	*GormRepository[models.BoardRole]
}

func NewBoardRoleRepository(db *gorm.DB) BoardRoleRepository {
	return &GormBoardRoleRepository{
		GormRepository: NewGormRepository[models.BoardRole](db),
	}
}

func (r *GormBoardRoleRepository) Migrate() error {
	return r.db.AutoMigrate(&models.BoardRole{})
}
