package database

import (
	"log"
	"server/config"
)

type Repository[T any] interface {
	Migrate() error
	GetAll() ([]T, error)
	GetByID(id uint) (T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
	Count() (int64, error)
}

type Database struct {
	UserRepository UserRepository
}

var DB Database

func Init() {
	initFuncs := map[string]func() (Database, error){
		"sqlite": NewSqlite,
		"mysql":  NewMySQL,
	}

	initFunc, exists := initFuncs[config.CFG.DBType]
	if !exists {
		log.Fatalf("Unsupported database type: %s", config.CFG.DBType)
	}

	database, err := initFunc()
	if err != nil {
		log.Fatalf("Failed to initialize %s database: %v", config.CFG.DBType, err)
	}

	err = database.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate %s database: %v", config.CFG.DBType, err)
	}

	DB = database
}
