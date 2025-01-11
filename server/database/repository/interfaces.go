package repository

import "gorm.io/gorm"

type QueryOption func(*gorm.DB) *gorm.DB

type Repository[T any] interface {
	Migrate() error
	GetAll(opts ...QueryOption) ([]T, error)
	GetFirst(opts ...QueryOption) (T, error)
	GetByID(id uint, opts ...QueryOption) (T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
	Count(opts ...QueryOption) (int64, error)
}

type Database struct {
	UserRepository          UserRepository
	BoardRepository         BoardRepository
	BoardRoleRepository     BoardRoleRepository
	SwimlaneRepository      SwimlaneRepository
	TaskRepository          TaskRepository
	UserBoardRoleRepository UserBoardRoleRepository
	CommentRepository       CommentRepository
	ColumnRepository        ColumnRepository
	FileRepository          FileRepository
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

	return nil
}
