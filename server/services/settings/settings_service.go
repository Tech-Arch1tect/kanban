package settings

import (
	"server/database/repository"
)

type SettingsService struct {
	db *repository.Database
}

func NewSettingsService(db *repository.Database) *SettingsService {
	return &SettingsService{db: db}
}
