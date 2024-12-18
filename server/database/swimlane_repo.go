package database

import (
	"server/models"
	"sort"

	"gorm.io/gorm"
)

type SwimlaneRepository interface {
	Repository[models.Swimlane]
	GetByBoardID(boardID uint) ([]models.Swimlane, error)
	GetNextOrder(boardID uint) (int, error)
}

type GormSwimlaneRepository struct {
	*GormRepository[models.Swimlane]
}

func NewSwimlaneRepository(db *gorm.DB) SwimlaneRepository {
	return &GormSwimlaneRepository{
		GormRepository: NewGormRepository[models.Swimlane](db),
	}
}

func (r *GormSwimlaneRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Swimlane{})
}

func (r *GormSwimlaneRepository) GetByBoardID(boardID uint) ([]models.Swimlane, error) {
	var swimlanes []models.Swimlane
	err := r.db.Where("board_id = ?", boardID).Find(&swimlanes).Error
	return r.SortByOrder(swimlanes), err
}

func (r *GormSwimlaneRepository) GetNextOrder(boardID uint) (int, error) {
	var swimlane models.Swimlane
	err := r.db.Where("board_id = ?", boardID).Order("`order` DESC").First(&swimlane).Error
	if err != nil {
		return 0, err
	}
	return swimlane.Order + 1, nil
}

func (r *GormSwimlaneRepository) SortByOrder(swimlanes []models.Swimlane) []models.Swimlane {
	sort.Slice(swimlanes, func(i, j int) bool { return swimlanes[i].Order < swimlanes[j].Order })
	return swimlanes
}
