package settings

import (
	"server/database/repository"
	"server/models"

	"gorm.io/gorm"
)

func (ss *SettingsService) GetSettings(userID uint) (models.Settings, error) {
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
