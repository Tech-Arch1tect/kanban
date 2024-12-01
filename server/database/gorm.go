package database

import (
	"server/models"

	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func (d *GormDB) GetUsers(page, pageSize int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := d.db.Model(&models.User{})

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

func (d *GormDB) GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := d.db.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *GormDB) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := d.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *GormDB) UpdateUserByID(id string, user models.User) error {
	if err := d.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (d *GormDB) DeleteUserByID(id string) error {
	if err := d.db.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (d *GormDB) CreateUser(user models.User) error {
	if err := d.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (d *GormDB) CountUsers() (int64, error) {
	var count int64
	if err := d.db.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
