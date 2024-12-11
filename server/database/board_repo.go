package database

import (
	"server/models"

	"gorm.io/gorm"
)

type BoardRepository interface {
	Repository[models.Board]
	GetWithPreload(id uint) (models.Board, error)
}

type GormBoardRepository struct {
	*GormRepository[models.Board]
}

func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &GormBoardRepository{
		GormRepository: NewGormRepository[models.Board](db),
	}
}

func (r *GormBoardRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Board{})
}

func (r *GormBoardRepository) GetWithPreload(id uint) (models.Board, error) {
	var board models.Board
	result := r.db.Preload("Permissions").Preload("Swimlanes").Preload("Tasks").Preload("Owner").First(&board, id)
	return board, result.Error
}
