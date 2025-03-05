package taskActivity

import (
	"server/config"
	"server/database/repository"
	"server/models"
	"server/services/eventBus"

	"go.uber.org/zap"
)

type TaskActivityService struct {
	db     *repository.Database
	config *config.Config
	te     *eventBus.EventBus[models.Task]
	ce     *eventBus.EventBus[models.Comment]
	fe     *eventBus.EventBus[models.File]
	le     *eventBus.EventBus[models.TaskLinks]
	lee    *eventBus.EventBus[models.TaskExternalLink]
	cre    *eventBus.EventBus[models.Reaction]
	logger *zap.Logger
}

func NewTaskActivityService(db *repository.Database, config *config.Config, te *eventBus.EventBus[models.Task], ce *eventBus.EventBus[models.Comment], fe *eventBus.EventBus[models.File], le *eventBus.EventBus[models.TaskLinks], lee *eventBus.EventBus[models.TaskExternalLink], cre *eventBus.EventBus[models.Reaction], logger *zap.Logger) *TaskActivityService {
	return &TaskActivityService{db: db, config: config, te: te, ce: ce, fe: fe, le: le, lee: lee, cre: cre, logger: logger}
}
