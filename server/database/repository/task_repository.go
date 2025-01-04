package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Repository[models.Task]
}

type GormTaskRepository struct {
	*GormRepository[models.Task]
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &GormTaskRepository{
		GormRepository: NewGormRepository[models.Task](db),
	}
}

func (r *GormTaskRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Task{})
}
