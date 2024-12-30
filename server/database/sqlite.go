package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSQLiteDB(cfg DBConfig) (*Database, error) {
	sqliteCfg := cfg.GetSQLiteConfig()
	gormDB, err := gorm.Open(sqlite.Open(sqliteCfg.GetFilePath()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{
		UserRepository: NewUserRepository(gormDB),
	}, nil
}
