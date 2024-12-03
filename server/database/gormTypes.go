package database

import (
	"errors"
	"fmt"
	"server/config"
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite() (Database, error) {
	db, err := gorm.Open(sqlite.Open(config.CFG.SQLite.FilePath), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return Database{}, err
	}

	userRepo := NewUserRepository(db)

	return Database{
		UserRepository: userRepo,
	}, nil
}

func NewMySQL() (Database, error) {
	if config.CFG.MySQL.User == "" || config.CFG.MySQL.Password == "" || config.CFG.MySQL.Host == "" || config.CFG.MySQL.Port == "" || config.CFG.MySQL.Database == "" {
		return Database{}, errors.New("missing mysql config")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.CFG.MySQL.User, config.CFG.MySQL.Password, config.CFG.MySQL.Host, config.CFG.MySQL.Port, config.CFG.MySQL.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return Database{}, err
	}

	userRepo := NewUserRepository(db)

	return Database{
		UserRepository: userRepo,
	}, nil
}
