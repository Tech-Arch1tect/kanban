package database

import (
	"server/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Repository[models.Comment]
	GetWithPreload(id uint) (models.Comment, error)
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

func (r *GormCommentRepository) GetWithPreload(id uint) (models.Comment, error) {
	var comment models.Comment
	result := r.db.Preload("User").First(&comment, id)
	return comment, result.Error
}
