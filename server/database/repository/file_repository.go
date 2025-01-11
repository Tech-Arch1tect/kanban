package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type FileRepository interface {
	Repository[models.File]
}

type GormFileRepository struct {
	*GormRepository[models.File]
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &GormFileRepository{
		GormRepository: NewGormRepository[models.File](db),
	}
}

func (r *GormFileRepository) Migrate() error {
	return r.db.AutoMigrate(&models.File{})
}
