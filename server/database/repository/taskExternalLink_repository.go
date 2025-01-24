package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type TaskExternalLinkRepository interface {
	Repository[models.TaskExternalLink]
}

type GormTaskExternalLinkRepository struct {
	*GormRepository[models.TaskExternalLink]
}

func NewTaskExternalLinkRepository(db *gorm.DB) TaskExternalLinkRepository {
	return &GormTaskExternalLinkRepository{
		GormRepository: NewGormRepository[models.TaskExternalLink](db),
	}
}

func (r *GormTaskExternalLinkRepository) Migrate() error {
	return r.db.AutoMigrate(&models.TaskExternalLink{})
}
