package database

import (
	"errors"
	"fmt"
	"server/config"
	"server/database/repository"
	"server/models"

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
		UserRepository:                      repository.NewGormRepository[models.User](gormDB),
		BoardRepository:                     repository.NewGormRepository[models.Board](gormDB),
		BoardRoleRepository:                 repository.NewGormRepository[models.BoardRole](gormDB),
		SwimlaneRepository:                  repository.NewGormRepository[models.Swimlane](gormDB),
		TaskRepository:                      repository.NewGormRepository[models.Task](gormDB),
		UserBoardRoleRepository:             repository.NewGormRepository[models.UserBoardRole](gormDB),
		CommentRepository:                   repository.NewGormRepository[models.Comment](gormDB),
		ColumnRepository:                    repository.NewGormRepository[models.Column](gormDB),
		FileRepository:                      repository.NewGormRepository[models.File](gormDB),
		BoardInviteRepository:               repository.NewGormRepository[models.BoardInvite](gormDB),
		TaskLinkRepository:                  repository.NewGormRepository[models.TaskLinks](gormDB),
		TaskExternalLinkRepository:          repository.NewGormRepository[models.TaskExternalLink](gormDB),
		SettingsRepository:                  repository.NewGormRepository[models.Settings](gormDB),
		NotificationConfigurationRepository: repository.NewNotificationConfigurationRepository(gormDB),
		NotificationEventRepository:         repository.NewGormRepository[models.NotificationEvent](gormDB),
		NotificationLogRepository:           repository.NewGormRepository[models.NotificationLog](gormDB),
		CommentReactionRepository:           repository.NewGormRepository[models.Reaction](gormDB),
		TaskActivityRepository:              repository.NewGormRepository[models.TaskActivity](gormDB),
	}, nil
}
