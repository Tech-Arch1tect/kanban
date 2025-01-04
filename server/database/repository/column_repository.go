package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type ColumnRepository interface {
	Repository[models.Column]
}

type GormColumnRepository struct {
	*GormRepository[models.Column]
}

func NewColumnRepository(db *gorm.DB) ColumnRepository {
	return &GormColumnRepository{
		GormRepository: NewGormRepository[models.Column](db),
	}
}

func (r *GormColumnRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Column{})
}
