package database

import (
	"server/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[models.User]
	GetByEmail(email string) (models.User, error)
	PaginatedSearch(page, pageSize int, search string) ([]models.User, int64, error)
}

type GormUserRepository struct {
	*GormRepository[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{
		GormRepository: NewGormRepository[models.User](db),
	}
}

func (r *GormUserRepository) Migrate() error {
	return r.db.AutoMigrate(&models.User{})
}

func (r *GormUserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *GormUserRepository) PaginatedSearch(page, pageSize int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.Model(&models.User{})

	if search != "" {
		query = query.Where("email LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order("created_at DESC")

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
