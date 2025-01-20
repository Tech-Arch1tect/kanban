package services

import (
	"server/database/repository"
	"server/models"

	"gorm.io/gorm"
)

type SettingsService struct {
	db *repository.Database
}

func NewSettingsService(db *repository.Database) *SettingsService {
	return &SettingsService{db: db}
}

func (ss *SettingsService) GetSettings(userID uint) (models.Settings, error) {
	var settings models.Settings
	settings, err := ss.db.SettingsRepository.GetFirst(repository.WithWhere("user_id = ?", userID))
	if err == gorm.ErrRecordNotFound {
		settings = models.Settings{
			UserID: userID,
		}
		err = ss.db.SettingsRepository.Create(&settings)
		if err != nil {
			return models.Settings{}, err
		}
		return settings, nil
	}
	if err != nil {
		return models.Settings{}, err
	}
	return settings, nil
}

func (ss *SettingsService) UpdateSettings(userID uint, settings models.Settings) error {
	oldSettings, err := ss.GetSettings(userID)
	if err != nil {
		return err
	}
	settings.ID = oldSettings.ID
	settings.UserID = oldSettings.UserID
	return ss.db.SettingsRepository.Update(&settings)
}
