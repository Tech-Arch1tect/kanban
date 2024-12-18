package database

import (
	"server/models"
	"sort"

	"gorm.io/gorm"
)

type ColumnRepository interface {
	Repository[models.Column]
	GetByBoardID(boardID uint) ([]models.Column, error)
	GetNextOrder(boardID uint) (int, error)
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

func (r *GormColumnRepository) GetNextOrder(boardID uint) (int, error) {
	var column models.Column
	err := r.db.Where("board_id = ?", boardID).Order("`order` DESC").First(&column).Error
	if err != nil {
		return 0, err
	}
	return column.Order + 1, nil
}

func (r *GormColumnRepository) SortByOrder(columns []models.Column) []models.Column {
	sort.Slice(columns, func(i, j int) bool { return columns[i].Order < columns[j].Order })
	return columns
}

func (r *GormColumnRepository) GetByBoardID(boardID uint) ([]models.Column, error) {
	var columns []models.Column
	err := r.db.Where("board_id = ?", boardID).Order("`order` ASC").Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return r.SortByOrder(columns), nil
}
