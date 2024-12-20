package database

import (
	"errors"
	"server/models"
	"slices"
	"strings"

	"gorm.io/gorm"
)

var queryMap = map[string]string{
	"status":      "status",      // =
	"title":       "title",       // LIKE %value%
	"description": "description", // LIKE %value%
	"board_id":    "board_id",    // =
	"swimlane_id": "swimlane_id", // =
}

type TaskRepository interface {
	Repository[models.Task]
	GetWithPreload(id uint) (models.Task, error)
	GetWithQuery(query string, user models.User) ([]models.Task, error)
	GetByColumnAndSwimlane(columnID uint, swimlaneID uint) ([]models.Task, error)
	GetPosition(columnID uint, swimlaneID uint) (int, error)
	RePositionAll(columnID uint, swimlaneID uint) error
}

type GormTaskRepository struct {
	*GormRepository[models.Task]
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &GormTaskRepository{
		GormRepository: NewGormRepository[models.Task](db),
	}
}

func (r *GormTaskRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Task{})
}

func (r *GormTaskRepository) GetWithPreload(id uint) (models.Task, error) {
	var task models.Task
	result := r.db.Preload("Swimlane").Preload("Column").Preload("Comments").Preload("Assignee").Preload("Creator").First(&task, id)
	return task, result.Error
}

func (r *GormTaskRepository) GetWithQuery(query string, user models.User) ([]models.Task, error) {
	permissions, err := DB.BoardPermissionRepository.GetPermissionsByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	boardIDs := make([]uint, 0, len(permissions))
	for _, p := range permissions {
		if p.GeneralAccess || p.Edit || p.Delete {
			boardIDs = append(boardIDs, p.BoardID)
		}
	}

	// find board id's where the user is owner
	ownerBoardIDs, err := DB.BoardRepository.GetAllByAccess(user.ID)
	if err != nil {
		return nil, err
	}

	for _, board := range ownerBoardIDs {
		if !slices.Contains(boardIDs, board.ID) {
			boardIDs = append(boardIDs, board.ID)
		}
	}

	if len(boardIDs) == 0 {
		return []models.Task{}, nil
	}

	q := TaskQuery{Query: query}
	filters, err := q.Parse()
	if err != nil {
		return nil, err
	}

	tx := r.db.Model(&models.Task{}).Where("board_id IN ?", boardIDs)

	for key, value := range filters {
		if column, ok := queryMap[key]; ok {
			switch key {
			case "title", "description":
				tx = tx.Where(column+" LIKE ?", "%"+value+"%")
			default:
				tx = tx.Where(column+" = ?", value)
			}
		}
	}

	var tasks []models.Task
	if err := tx.Order("position ASC").Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

type TaskQuery struct {
	Query string `json:"query"`
}

// Parse turns a simple query like "status:open title:\"Some Title\"" into a map.
// Example: "status:open title:test" -> {"status": "open", "title": "test"}
func (q *TaskQuery) Parse() (map[string]string, error) {
	filters := make(map[string]string)
	q.Query = strings.TrimSpace(q.Query)
	if q.Query == "" {
		return filters, nil
	}

	parts := strings.Fields(q.Query)
	for _, part := range parts {
		pair := strings.SplitN(part, ":", 2)
		if len(pair) != 2 {
			return nil, errors.New("invalid query format")
		}
		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])
		if key == "" || value == "" {
			return nil, errors.New("invalid query format")
		}
		filters[key] = value
	}

	return filters, nil
}

func (r *GormTaskRepository) GetByColumnAndSwimlane(columnID uint, swimlaneID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("column_id = ? AND swimlane_id = ?", columnID, swimlaneID).Order("position ASC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *GormTaskRepository) GetPosition(columnID uint, swimlaneID uint) (int, error) {
	var task models.Task
	if err := r.db.Where("column_id = ? AND swimlane_id = ?", columnID, swimlaneID).Order("position DESC").First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return task.Position, nil
}

func (r *GormTaskRepository) RePositionAll(columnID uint, swimlaneID uint) error {
	tasks, err := r.GetByColumnAndSwimlane(columnID, swimlaneID)
	if err != nil {
		return err
	}
	for i, task := range tasks {
		task.Position = i
		if err := r.Update(&task); err != nil {
			return err
		}
	}
	return nil
}
