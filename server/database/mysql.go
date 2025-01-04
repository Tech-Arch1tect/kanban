package database

import (
	"errors"
	"fmt"
	"server/config"
	"server/database/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMySQLDB(cfg *config.Config) (*repository.Database, error) {
	mysqlCfg := cfg.Database.MySQL

	if mysqlCfg.User == "" ||
		mysqlCfg.Password == "" ||
		mysqlCfg.Host == "" ||
		mysqlCfg.Port == "" ||
		mysqlCfg.Database == "" {
		return nil, errors.New("missing MySQL config")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Database,
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &repository.Database{
		UserRepository:                repository.NewUserRepository(gormDB),
		BoardRepository:               repository.NewBoardRepository(gormDB),
		SwimlaneRepository:            repository.NewSwimlaneRepository(gormDB),
		TaskRepository:                repository.NewTaskRepository(gormDB),
		UserBoardPermissionRepository: repository.NewUserBoardPermissionRepository(gormDB),
		CommentRepository:             repository.NewCommentRepository(gormDB),
		ColumnRepository:              repository.NewColumnRepository(gormDB),
	}, nil
}
