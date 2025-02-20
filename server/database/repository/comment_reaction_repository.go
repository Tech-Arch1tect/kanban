package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type CommentReactionRepository interface {
	Repository[models.Reaction]
}

type GormCommentReactionRepository struct {
	*GormRepository[models.Reaction]
}

func NewCommentReactionRepository(db *gorm.DB) CommentReactionRepository {
	return &GormCommentReactionRepository{
		GormRepository: NewGormRepository[models.Reaction](db),
	}
}

func (r *GormCommentReactionRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Reaction{})
}
