package database

import (
	"errors"
	"server/models"

	"gorm.io/gorm"
)

type BoardRepository interface {
	Repository[models.Board]
	GetWithPreload(id uint) (models.Board, error)
	GetAllByAccess(userID uint) ([]models.Board, error)
	GetPermission(userID uint, boardID uint) (models.BoardPermission, error)
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
	result := r.db.Preload("Permissions").Preload("Swimlanes").Preload("Columns").Preload("Owner").First(&board, id)
	return board, result.Error
}

func (r *GormBoardRepository) GetAllByAccess(userID uint) ([]models.Board, error) {
	var boards []models.Board
	result := r.db.Preload("Permissions").Preload("Swimlanes").Preload("Columns").Preload("Owner").Joins("LEFT JOIN board_permissions ON boards.id = board_permissions.board_id").Where("board_permissions.user_id = ?", userID).Or("boards.owner_id = ?", userID).Find(&boards)
	return boards, result.Error
}

func (r *GormBoardRepository) GetPermission(userID uint, boardID uint) (models.BoardPermission, error) {
	var board models.Board
	result := r.db.Where("id = ?", boardID).Preload("Permissions").First(&board)
	if result.Error != nil {
		return models.BoardPermission{}, result.Error
	}

	for _, permission := range board.Permissions {
		if permission.UserID == userID {
			return permission, nil
		}
	}
	return models.BoardPermission{}, errors.New("no permission found")
}
