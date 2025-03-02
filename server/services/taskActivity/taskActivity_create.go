package taskActivity

import "server/models"

func (s *TaskActivityService) CreateTaskActivity(taskID uint, userID uint, event string, oldData string, newData string) error {

	taskActivity := &models.TaskActivity{
		TaskID:  taskID,
		UserID:  userID,
		Event:   event,
		OldData: oldData,
		NewData: newData,
	}
	return s.db.TaskActivityRepository.Create(taskActivity)
}
