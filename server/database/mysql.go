package database

import (
	"errors"
	"fmt"
	"server/database/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMySQLDB(cfg DBConfig) (*repository.Database, error) {
	mysqlCfg := cfg.GetMySQLConfig()

	if mysqlCfg.GetUser() == "" ||
		mysqlCfg.GetPassword() == "" ||
		mysqlCfg.GetHost() == "" ||
		mysqlCfg.GetPort() == "" ||
		mysqlCfg.GetDatabase() == "" {
		return nil, errors.New("missing MySQL config")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.GetUser(),
		mysqlCfg.GetPassword(),
		mysqlCfg.GetHost(),
		mysqlCfg.GetPort(),
		mysqlCfg.GetDatabase(),
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &repository.Database{
		UserRepository: repository.NewUserRepository(gormDB),
	}, nil
}
