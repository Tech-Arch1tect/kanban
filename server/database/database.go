package database

import (
	"fmt"
	"server/config"
	"server/database/repository"
)

func Init(cfg *config.Config) (repository.Database, error) {
	initFuncs := map[string]func(*config.Config) (*repository.Database, error){
		"sqlite": newSQLiteDB,
		"mysql":  newMySQLDB,
	}

	initFunc, exists := initFuncs[cfg.Database.Type]
	if !exists {
		return repository.Database{}, fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}

	db, err := initFunc(cfg)
	if err != nil {
		return repository.Database{}, fmt.Errorf("failed to initialise %s database: %w", cfg.Database.Type, err)
	}

	if err := db.Migrate(); err != nil {
		return repository.Database{}, fmt.Errorf("failed to migrate %s database: %w", cfg.Database.Type, err)
	}

	return *db, nil
}
