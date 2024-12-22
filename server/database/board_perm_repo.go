package database

import (
	"server/models"
	"slices"

	"gorm.io/gorm"
)

type BoardPermissionRepository interface {
	Repository[models.BoardPermission]
	GetPermissionsByUserID(userID uint) ([]models.BoardPermission, error)
	GetUsersWithAccessToBoard(boardID uint) ([]models.User, error)
}

type GormBoardPermissionRepository struct {
	*GormRepository[models.BoardPermission]
}

func NewBoardPermissionRepository(db *gorm.DB) BoardPermissionRepository {
	return &GormBoardPermissionRepository{
		GormRepository: NewGormRepository[models.BoardPermission](db),
	}
}

func (r *GormBoardPermissionRepository) Migrate() error {
	return r.db.AutoMigrate(&models.BoardPermission{})
}

func (r *GormBoardPermissionRepository) GetPermissionsByUserID(userID uint) ([]models.BoardPermission, error) {
	var permissions []models.BoardPermission
	result := r.db.Where("user_id = ?", userID).Find(&permissions)
	return permissions, result.Error
}

func (r *GormBoardPermissionRepository) GetUsersWithAccessToBoard(boardID uint) ([]models.User, error) {
	var permissions []models.BoardPermission
	result := r.db.Where("board_id = ?", boardID).Find(&permissions).Preload("User")
	if result.Error != nil {
		return nil, result.Error
	}

	var users []models.User
	for _, permission := range permissions {
		users = append(users, permission.User)
	}

	admins, err := DB.UserRepository.GetUsersByRole(models.RoleAdmin)
	if err != nil {
		return nil, err
	}
	for _, admin := range admins {
		if !slices.Contains(users, admin) {
			users = append(users, admin)
		}
	}
	return users, nil
}
