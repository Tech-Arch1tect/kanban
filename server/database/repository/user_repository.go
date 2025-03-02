package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[models.User]
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
