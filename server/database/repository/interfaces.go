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
	UserRepository UserRepository
}

func (db *Database) Migrate() error {
	return db.UserRepository.Migrate()
}
