package database

import (
	"server/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Repository[models.Task]
	GetWithPreload(id uint) (models.Task, error)
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

func (r *GormTaskRepository) GetWithPreload(id uint) (models.Task, error) {
	var task models.Task
	result := r.db.Preload("Swimlane").Preload("Comments").First(&task, id)
	return task, result.Error
}
