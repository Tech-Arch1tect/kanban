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
	r := &GormDB{}
	db, err := gorm.Open(sqlite.Open(config.CFG.SQLite.FilePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	r.db = db
	db.AutoMigrate(&models.User{})
	return r, nil
}

func NewMySQL() (Database, error) {
	r := &GormDB{}
	if config.CFG.MySQL.User == "" || config.CFG.MySQL.Password == "" || config.CFG.MySQL.Host == "" || config.CFG.MySQL.Port == "" || config.CFG.MySQL.Database == "" {
		return nil, errors.New("missing mysql config")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.CFG.MySQL.User, config.CFG.MySQL.Password, config.CFG.MySQL.Host, config.CFG.MySQL.Port, config.CFG.MySQL.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	r.db = db
	db.AutoMigrate(&models.User{})
	return r, nil
}
