package database

import (
	"log"
)

type DBConfig interface {
	GetDBType() string
	GetSQLiteConfig() SQLiteConfig
	GetMySQLConfig() MySQLConfig
}

type SQLiteConfig interface {
	GetFilePath() string
}

type MySQLConfig interface {
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetDatabase() string
}

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

func Init(cfg DBConfig) {
	initFuncs := map[string]func(DBConfig) (Database, error){
		"sqlite": NewSqlite,
		"mysql":  NewMySQL,
	}

	initFunc, exists := initFuncs[cfg.GetDBType()]
	if !exists {
		log.Fatalf("Unsupported database type: %s", cfg.GetDBType())
	}

	database, err := initFunc(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize %s database: %v", cfg.GetDBType(), err)
	}

	err = database.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate %s database: %v", cfg.GetDBType(), err)
	}

	DB = database
}
