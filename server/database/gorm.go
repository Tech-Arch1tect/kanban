package database

import (
	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) GetAll() ([]T, error) {
	var entities []T
	result := r.db.Find(&entities)
	return entities, result.Error
}

func (r *GormRepository[T]) GetByID(id uint) (T, error) {
	var entity T
	result := r.db.First(&entity, id)
	return entity, result.Error
}

func (r *GormRepository[T]) Create(entity *T) error {
	result := r.db.Create(entity)
	return result.Error
}

func (r *GormRepository[T]) Update(entity *T) error {
	result := r.db.Save(entity)
	return result.Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	result := r.db.Delete(new(T), id)
	return result.Error
}

func (r *GormRepository[T]) Count() (int64, error) {
	var count int64
	result := r.db.Model(new(T)).Count(&count)
	return count, result.Error
}
