package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Repository[models.Comment]
}

type GormCommentRepository struct {
	*GormRepository[models.Comment]
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &GormCommentRepository{
		GormRepository: NewGormRepository[models.Comment](db),
	}
}

func (r *GormCommentRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Comment{})
}
