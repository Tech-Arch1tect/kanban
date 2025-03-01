package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type TaskActivityRepository interface {
	Repository[models.TaskActivity]
}

type GormTaskActivityRepository struct {
	*GormRepository[models.TaskActivity]
}

func NewTaskActivityRepository(db *gorm.DB) TaskActivityRepository {
	return &GormTaskActivityRepository{
		GormRepository: NewGormRepository[models.TaskActivity](db),
	}
}

func (r *GormTaskActivityRepository) Migrate() error {
	return r.db.AutoMigrate(&models.TaskActivity{})
}
