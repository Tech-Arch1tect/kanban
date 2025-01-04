package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type BoardRepository interface {
	Repository[models.Board]
}

type GormBoardRepository struct {
	*GormRepository[models.Board]
}

func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &GormBoardRepository{
		GormRepository: NewGormRepository[models.Board](db),
	}
}

func (r *GormBoardRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Board{})
}
