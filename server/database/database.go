package database

import (
	"fmt"
)

type Database struct {
	UserRepository UserRepository
}

var DB Database

func Init(cfg DBConfig) error {
	initFuncs := map[string]func(DBConfig) (*Database, error){
		"sqlite": newSQLiteDB,
		"mysql":  newMySQLDB,
	}

	initFunc, exists := initFuncs[cfg.GetDBType()]
	if !exists {
		return fmt.Errorf("unsupported database type: %s", cfg.GetDBType())
	}

	db, err := initFunc(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize %s database: %w", cfg.GetDBType(), err)
	}

	if err := db.Migrate(); err != nil {
		return fmt.Errorf("failed to migrate %s database: %w", cfg.GetDBType(), err)
	}

	DB = *db
	return nil
}

func (db *Database) Migrate() error {
	return db.UserRepository.Migrate()
}
