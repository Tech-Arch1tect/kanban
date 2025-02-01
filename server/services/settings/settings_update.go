package settings

import "server/models"

func (ss *SettingsService) UpdateSettings(userID uint, settings models.Settings) error {
	oldSettings, err := ss.GetSettings(userID)
	if err != nil {
		return err
	}
	settings.ID = oldSettings.ID
	settings.UserID = oldSettings.UserID
	return ss.db.SettingsRepository.Update(&settings)
}
