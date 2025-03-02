package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type QueryOption func(*gorm.DB) *gorm.DB

type Repository[T any] interface {
	Migrate() error
	GetAll(opts ...QueryOption) ([]T, error)
	GetFirst(opts ...QueryOption) (T, error)
	GetByID(id uint, opts ...QueryOption) (T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
	HardDelete(id uint) error
	Count(opts ...QueryOption) (int64, error)
	PaginatedSearch(page, pageSize int, search string, searchField string, orderBy string, opts ...QueryOption) ([]T, int64, error)
}

type Database struct {
	UserRepository                      *GormRepository[models.User]
	BoardRepository                     *GormRepository[models.Board]
	BoardRoleRepository                 *GormRepository[models.BoardRole]
	SwimlaneRepository                  *GormRepository[models.Swimlane]
	TaskRepository                      *GormRepository[models.Task]
	UserBoardRoleRepository             *GormRepository[models.UserBoardRole]
	CommentRepository                   *GormRepository[models.Comment]
	ColumnRepository                    *GormRepository[models.Column]
	FileRepository                      *GormRepository[models.File]
	BoardInviteRepository               *GormRepository[models.BoardInvite]
	TaskLinkRepository                  *GormRepository[models.TaskLinks]
	TaskExternalLinkRepository          *GormRepository[models.TaskExternalLink]
	SettingsRepository                  *GormRepository[models.Settings]
	NotificationConfigurationRepository NotificationConfigurationRepository
	NotificationEventRepository         *GormRepository[models.NotificationEvent]
	NotificationLogRepository           *GormRepository[models.NotificationLog]
	CommentReactionRepository           *GormRepository[models.Reaction]
	TaskActivityRepository              *GormRepository[models.TaskActivity]
}

func (db *Database) Migrate() error {
	err := db.UserRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.BoardRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.BoardRoleRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.SwimlaneRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.TaskRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.UserBoardRoleRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.CommentRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.ColumnRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.FileRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.BoardInviteRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.TaskLinkRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.SettingsRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.TaskExternalLinkRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.NotificationConfigurationRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.NotificationEventRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.NotificationLogRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.CommentReactionRepository.Migrate()
	if err != nil {
		return err
	}

	err = db.TaskActivityRepository.Migrate()
	if err != nil {
		return err
	}

	return nil
}
