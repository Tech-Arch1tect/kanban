package database

import (
	"errors"
	"fmt"
	"server/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite() (Database, error) {
	db, err := gorm.Open(sqlite.Open(config.CFG.SQLite.FilePath), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	userRepo := NewUserRepository(db)
	boardRepo := NewBoardRepository(db)
	boardPermRepo := NewBoardPermissionRepository(db)
	taskRepo := NewTaskRepository(db)
	swimlaneRepo := NewSwimlaneRepository(db)
	commentRepo := NewCommentRepository(db)
	columnRepo := NewColumnRepository(db)
	return Database{
		UserRepository:            userRepo,
		BoardRepository:           boardRepo,
		BoardPermissionRepository: boardPermRepo,
		TaskRepository:            taskRepo,
		SwimlaneRepository:        swimlaneRepo,
		CommentRepository:         commentRepo,
		ColumnRepository:          columnRepo,
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

	userRepo := NewUserRepository(db)
	boardRepo := NewBoardRepository(db)
	boardPermRepo := NewBoardPermissionRepository(db)
	taskRepo := NewTaskRepository(db)
	swimlaneRepo := NewSwimlaneRepository(db)
	commentRepo := NewCommentRepository(db)
	columnRepo := NewColumnRepository(db)
	return Database{
		UserRepository:            userRepo,
		BoardRepository:           boardRepo,
		BoardPermissionRepository: boardPermRepo,
		TaskRepository:            taskRepo,
		CommentRepository:         commentRepo,
		SwimlaneRepository:        swimlaneRepo,
		ColumnRepository:          columnRepo,
	}, nil
}

func (d *Database) Migrate() error {
	// migrate all models
	err := d.UserRepository.Migrate()
	if err != nil {
		return err
	}

	err = d.BoardRepository.Migrate()
	if err != nil {
		return err
	}

	err = d.BoardPermissionRepository.Migrate()
	if err != nil {
		return err
	}

	err = d.TaskRepository.Migrate()
	if err != nil {
		return err
	}

	err = d.SwimlaneRepository.Migrate()
	if err != nil {
		return err
	}

	err = d.CommentRepository.Migrate()
	if err != nil {
		return err
	}

	err = d.ColumnRepository.Migrate()
	if err != nil {
		return err
	}

	return nil
}
