package taskActivity

import (
	"server/config"
	"server/database/repository"
	"server/models"
	"server/services/eventBus"
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
}

func NewTaskActivityService(db *repository.Database, config *config.Config, te *eventBus.EventBus[models.Task], ce *eventBus.EventBus[models.Comment], fe *eventBus.EventBus[models.File], le *eventBus.EventBus[models.TaskLinks], lee *eventBus.EventBus[models.TaskExternalLink], cre *eventBus.EventBus[models.Reaction]) *TaskActivityService {
	return &TaskActivityService{db: db, config: config, te: te, ce: ce, fe: fe, le: le, lee: lee, cre: cre}
}
