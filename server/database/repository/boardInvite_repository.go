package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type BoardInviteRepository interface {
	Repository[models.BoardInvite]
}

type GormBoardInviteRepository struct {
	*GormRepository[models.BoardInvite]
}

func NewBoardInviteRepository(db *gorm.DB) BoardInviteRepository {
	return &GormBoardInviteRepository{
		GormRepository: NewGormRepository[models.BoardInvite](db),
	}
}

func (r *GormBoardInviteRepository) Migrate() error {
	return r.db.AutoMigrate(&models.BoardInvite{})
}
