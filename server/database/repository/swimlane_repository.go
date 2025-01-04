package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type SwimlaneRepository interface {
	Repository[models.Swimlane]
}

type GormSwimlaneRepository struct {
	*GormRepository[models.Swimlane]
}

func NewSwimlaneRepository(db *gorm.DB) SwimlaneRepository {
	return &GormSwimlaneRepository{
		GormRepository: NewGormRepository[models.Swimlane](db),
	}
}

func (r *GormSwimlaneRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Swimlane{})
}
