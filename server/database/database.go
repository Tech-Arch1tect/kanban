package database

import (
	"log"
	"server/config"
	"server/models"
)

type Database interface {
	GetUsers(page, pageSize int, search string) ([]models.User, int64, error)
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUserByID(id string, user models.User) error
	DeleteUserByID(id string) error
	CreateUser(user models.User) error
	CountUsers() (int64, error)
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

	DB = database
}
