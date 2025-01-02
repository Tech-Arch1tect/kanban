package database

import (
	"server/config"
	"server/database/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSQLiteDB(cfg *config.Config) (*repository.Database, error) {
	sqliteCfg := cfg.Database.SQLite
	gormDB, err := gorm.Open(sqlite.Open(sqliteCfg.FilePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &repository.Database{
		UserRepository: repository.NewUserRepository(gormDB),
	}, nil
}
