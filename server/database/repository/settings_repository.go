package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type SettingsRepository interface {
	Repository[models.Settings]
}

type GormSettingsRepository struct {
	*GormRepository[models.Settings]
}

func NewSettingsRepository(db *gorm.DB) SettingsRepository {
	return &GormSettingsRepository{
		GormRepository: NewGormRepository[models.Settings](db),
	}
}

func (r *GormSettingsRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Settings{})
}
