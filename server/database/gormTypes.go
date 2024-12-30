package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite(cfg DBConfig) (Database, error) {
	sqliteCfg := cfg.GetSQLiteConfig()
	db, err := gorm.Open(sqlite.Open(sqliteCfg.GetFilePath()), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	userRepo := NewUserRepository(db)

	return Database{
		UserRepository: userRepo,
	}, nil
}

func NewMySQL(cfg DBConfig) (Database, error) {
	mysqlCfg := cfg.GetMySQLConfig()
	if mysqlCfg.GetUser() == "" || mysqlCfg.GetPassword() == "" || mysqlCfg.GetHost() == "" || mysqlCfg.GetPort() == "" || mysqlCfg.GetDatabase() == "" {
		return Database{}, errors.New("missing mysql config")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.GetUser(),
		mysqlCfg.GetPassword(),
		mysqlCfg.GetHost(),
		mysqlCfg.GetPort(),
		mysqlCfg.GetDatabase())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	userRepo := NewUserRepository(db)

	return Database{
		UserRepository: userRepo,
	}, nil
}

func (d *Database) Migrate() error {
	return d.UserRepository.Migrate()
}
