package repository

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) applyOptions(opts ...QueryOption) *gorm.DB {
	db := r.db
	for _, opt := range opts {
		db = opt(db)
	}
	return db
}

func (r *GormRepository[T]) GetAll(opts ...QueryOption) ([]T, error) {
	var entities []T
	db := r.applyOptions(opts...)
	result := db.Find(&entities)
	return entities, result.Error
}

func (r *GormRepository[T]) GetFirst(opts ...QueryOption) (T, error) {
	var entity T
	db := r.applyOptions(opts...)
	result := db.First(&entity)
	return entity, result.Error
}

func (r *GormRepository[T]) GetByID(id uint, opts ...QueryOption) (T, error) {
	var entity T
	db := r.applyOptions(opts...)
	result := db.First(&entity, id)
	return entity, result.Error
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GormRepository[T]) Update(entity *T) error {
	return r.db.Omit(clause.Associations).Save(entity).Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}

func (r *GormRepository[T]) HardDelete(id uint) error {
	return r.db.Unscoped().Delete(new(T), id).Error
}

func (r *GormRepository[T]) Count(opts ...QueryOption) (int64, error) {
	var count int64
	db := r.applyOptions(opts...)
	err := db.Model(new(T)).Count(&count).Error
	return count, err
}

func (r *GormRepository[T]) PaginatedSearch(page, pageSize int, search string, searchField string, orderBy string, opts ...QueryOption) ([]T, int64, error) {
	var (
		entities []T
		total    int64
	)

	db := r.applyOptions(opts...)
	if search != "" && searchField != "" {
		db = db.Where(fmt.Sprintf("%s LIKE ?", searchField), "%"+search+"%")
	}

	if err := db.Model(new(T)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if orderBy != "" {
		db = db.Order(orderBy)
	}

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		db = db.Offset(offset).Limit(pageSize)
	}

	if err := db.Find(&entities).Error; err != nil {
		return nil, 0, err
	}
	return entities, total, nil
}

func WithPreload(relations ...string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		for _, relation := range relations {
			db = db.Preload(relation)
		}
		return db
	}
}

func WithWhere(query interface{}, args ...interface{}) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func WithOrder(order string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithLimit(limit int) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func WithOrWhere(query string, args ...interface{}) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Or(query, args...)
	}
}

func WithCustom(f func(*gorm.DB) *gorm.DB) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return f(db)
	}
}

func WithJoin(join string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(join)
	}
}

func (r *GormRepository[T]) Migrate() error {
	return r.db.AutoMigrate(new(T))
}
