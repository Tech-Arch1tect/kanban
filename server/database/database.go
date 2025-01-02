package database

import (
	"fmt"
	"server/config"
	"server/database/repository"
)

var DB repository.Database

func Init(cfg *config.Config) error {
	initFuncs := map[string]func(*config.Config) (*repository.Database, error){
		"sqlite": newSQLiteDB,
		"mysql":  newMySQLDB,
	}

	initFunc, exists := initFuncs[cfg.Database.Type]
	if !exists {
		return fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}

	db, err := initFunc(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialise %s database: %w", cfg.Database.Type, err)
	}

	if err := db.Migrate(); err != nil {
		return fmt.Errorf("failed to migrate %s database: %w", cfg.Database.Type, err)
	}

	DB = *db
	return nil
}
