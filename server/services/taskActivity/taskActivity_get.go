package taskActivity

import (
	"server/database/repository"
	"server/models"
)

func (s *TaskActivityService) GetPaginatedTaskActivities(page, pageSize int) (taskActivities []models.TaskActivity, totalRecords int, totalPages int, err error) {
	taskActivities, totalRecordsint64, err := s.db.TaskActivityRepository.PaginatedSearch(page, pageSize, "", "", "created_at DESC", repository.WithPreload("User"))
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages = (int(totalRecordsint64) + pageSize - 1) / pageSize
	return taskActivities, int(totalRecordsint64), totalPages, nil
}
