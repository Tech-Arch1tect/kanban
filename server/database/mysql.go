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
		UserRepository:                      repository.NewUserRepository(gormDB),
		BoardRepository:                     repository.NewBoardRepository(gormDB),
		BoardRoleRepository:                 repository.NewBoardRoleRepository(gormDB),
		SwimlaneRepository:                  repository.NewSwimlaneRepository(gormDB),
		TaskRepository:                      repository.NewTaskRepository(gormDB),
		UserBoardRoleRepository:             repository.NewUserBoardRoleRepository(gormDB),
		CommentRepository:                   repository.NewCommentRepository(gormDB),
		ColumnRepository:                    repository.NewColumnRepository(gormDB),
		FileRepository:                      repository.NewFileRepository(gormDB),
		BoardInviteRepository:               repository.NewBoardInviteRepository(gormDB),
		TaskLinkRepository:                  repository.NewTaskLinkRepository(gormDB),
		TaskExternalLinkRepository:          repository.NewTaskExternalLinkRepository(gormDB),
		SettingsRepository:                  repository.NewSettingsRepository(gormDB),
		NotificationConfigurationRepository: repository.NewNotificationConfigurationRepository(gormDB),
		NotificationEventRepository:         repository.NewNotificationEventRepository(gormDB),
		NotificationLogRepository:           repository.NewNotificationLogRepository(gormDB),
		CommentReactionRepository:           repository.NewCommentReactionRepository(gormDB),
	}, nil
}
