package database

import (
	"server/database/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSQLiteDB(cfg DBConfig) (*repository.Database, error) {
	sqliteCfg := cfg.GetSQLiteConfig()
	gormDB, err := gorm.Open(sqlite.Open(sqliteCfg.GetFilePath()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &repository.Database{
		UserRepository: repository.NewUserRepository(gormDB),
	}, nil
}
