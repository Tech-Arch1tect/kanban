package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"server/config"
	"server/database/repository"
	"server/models"
	"sort"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskService struct {
	db     *repository.Database
	rs     *RoleService
	config *config.Config
}

func NewTaskService(db *repository.Database, rs *RoleService, config *config.Config) *TaskService {
	return &TaskService{db: db, rs: rs, config: config}
}

type CreateTaskRequest struct {
	ParentTaskID *uint  `json:"parent_task_id"`
	BoardID      uint   `json:"board_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	SwimlaneID   uint   `json:"swimlane_id"`
	ColumnID     uint   `json:"column_id"`
	Status       string `json:"status"`
	AssigneeID   uint   `json:"assignee_id"`
}

func (ts *TaskService) CreateTask(userID uint, request CreateTaskRequest) (models.Task, error) {
	can, _ := ts.rs.CheckRole(userID, request.BoardID, MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	_, err := ts.db.SwimlaneRepository.GetByID(request.SwimlaneID)
	if err != nil {
		return models.Task{}, err
	}

	_, err = ts.db.ColumnRepository.GetByID(request.ColumnID)
	if err != nil {
		return models.Task{}, err
	}

	taskPosition, err := ts.db.TaskRepository.GetFirst(
		repository.WithWhere("column_id = ? AND swimlane_id = ?", request.ColumnID, request.SwimlaneID),
		repository.WithOrder("position DESC"),
	)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Task{}, err
	}

	var assignee models.User
	if request.AssigneeID != 0 {
		assignee, err = ts.db.UserRepository.GetByID(request.AssigneeID)
		if err != nil {
			return models.Task{}, err
		}
	}

	task := models.Task{
		ParentTaskID: request.ParentTaskID,
		BoardID:      request.BoardID,
		Title:        request.Title,
		Description:  request.Description,
		SwimlaneID:   request.SwimlaneID,
		Status:       request.Status,
		ColumnID:     request.ColumnID,
		Position:     taskPosition.Position + 1,
		CreatorID:    userID,
		AssigneeID:   assignee.ID,
	}

	err = task.Validate()
	if err != nil {
		return models.Task{}, err
	}

	err = ts.db.TaskRepository.Create(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) DeleteTask(userID, taskID uint) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	err = ts.db.TaskRepository.Delete(task.ID)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) CommonPreloadFields() []string {
	return []string{"Board", "Swimlane", "Column", "Creator", "Assignee", "Comments", "Comments.User", "Files", "SrcLinks", "DstLinks", "SrcLinks.SrcTask", "SrcLinks.DstTask", "DstLinks.DstTask", "DstLinks.SrcTask", "ExternalLinks", "ParentTask",
		"Subtasks", "Subtasks.Board", "Subtasks.Swimlane", "Subtasks.Column", "Subtasks.Creator", "Subtasks.Assignee", "Subtasks.Comments", "Subtasks.Comments.User", "Subtasks.Files", "Subtasks.SrcLinks", "Subtasks.DstLinks", "Subtasks.SrcLinks.SrcTask", "Subtasks.SrcLinks.DstTask", "Subtasks.DstLinks.DstTask", "Subtasks.DstLinks.SrcTask", "Subtasks.ExternalLinks", "Subtasks.ParentTask"}
}

func (ts *TaskService) GetTask(userID, taskID uint) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetFirst(
		repository.WithWhere("id = ?", taskID),
		repository.WithPreload(ts.CommonPreloadFields()...),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	// inverse the dst links
	newDstLinks := []models.TaskLinks{}
	for _, link := range task.DstLinks {
		linkType := repository.InverseLinkTypeMap[link.LinkType]
		newDstLinks = append(newDstLinks, models.TaskLinks{
			Model: models.Model{
				ID: link.ID,
			},
			SrcTaskID: link.DstTaskID,
			SrcTask:   link.DstTask,
			DstTaskID: link.SrcTaskID,
			DstTask:   link.SrcTask,
			LinkType:  string(linkType),
		})
	}

	task.DstLinks = newDstLinks

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	return task, nil
}

type MoveTaskRequest struct {
	TaskID     uint    `json:"task_id"`
	ColumnID   uint    `json:"column_id"`
	SwimlaneID uint    `json:"swimlane_id"`
	Position   float64 `json:"position"`
}

func (ts *TaskService) MoveTask(userID uint, request MoveTaskRequest) (models.Task, error) {
	task, err := ts.GetTask(userID, request.TaskID)
	if err != nil {
		return models.Task{}, err
	}

	column, err := ts.db.ColumnRepository.GetByID(request.ColumnID)
	if err != nil {
		return models.Task{}, err
	}
	swimlane, err := ts.db.SwimlaneRepository.GetByID(request.SwimlaneID)
	if err != nil {
		return models.Task{}, err
	}
	if column.BoardID != swimlane.BoardID || column.BoardID != task.BoardID {
		return models.Task{}, errors.New("column, swimlane, and task must belong to the same board")
	}

	tasks, err := ts.db.TaskRepository.GetAll(
		repository.WithWhere("column_id = ? AND swimlane_id = ?", request.ColumnID, request.SwimlaneID),
		repository.WithOrder("position ASC"),
	)
	if err != nil {
		return models.Task{}, err
	}

	var filtered []models.Task
	for _, t := range tasks {
		if t.ID != task.ID {
			filtered = append(filtered, t)
		}
	}
	tasks = filtered

	var newPos float64

	if len(tasks) == 0 {
		newPos = request.Position
	} else {
		var nextPos float64
		foundNext := false
		for _, t := range tasks {
			if t.Position > request.Position {
				nextPos = t.Position
				foundNext = true
				break
			}
		}
		if !foundNext {
			newPos = request.Position + 1.0
		} else {
			newPos = (request.Position + nextPos) / 2.0
		}
	}

	task.Position = newPos
	task.ColumnID = request.ColumnID
	task.SwimlaneID = request.SwimlaneID

	if err := ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) GetTasksWithQuery(userID uint, boardID uint, query string) ([]models.Task, error) {
	can, _ := ts.rs.CheckRole(userID, boardID, ReaderRole)
	if !can {
		return nil, errors.New("forbidden")
	}

	statuses, assigneeEmail, searchTerm := parseQuery(query)

	var qopts []repository.QueryOption

	qopts = append(qopts, repository.WithWhere("board_id = ?", boardID))
	qopts = append(qopts, repository.WithWhere("parent_task_id IS NULL"))

	if len(statuses) > 0 {
		qopts = append(qopts, repository.WithWhere("status IN ?", statuses))
	}

	if assigneeEmail != "" {
		user, err := ts.db.UserRepository.GetFirst(repository.WithWhere("email = ?", assigneeEmail))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return []models.Task{}, nil
			}
			return nil, err
		}
		qopts = append(qopts, repository.WithWhere("assignee_id = ?", user.ID))
	}

	if searchTerm != "" {
		likeValue := fmt.Sprintf("%%%s%%", searchTerm)
		qopts = append(qopts, repository.WithCustom(func(db *gorm.DB) *gorm.DB {
			return db.Where("title LIKE ? OR description LIKE ?", likeValue, likeValue)
		}))
	}

	qopts = append(qopts, repository.WithPreload("Assignee", "Subtasks", "Subtasks.Assignee"))

	tasks, err := ts.db.TaskRepository.GetAll(qopts...)
	if err != nil {
		return nil, err
	}

	ts.SortTasks(tasks)

	return tasks, nil
}

func (ts *TaskService) SortTasks(tasks []models.Task) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Position < tasks[j].Position })
}

// parseQuery is a naive parser that extracts:
//   - `status:open|closed` → []string{"open","closed"}
//   - `assignee:email`  → "email"
//   - Everything else → appended into one search string
func parseQuery(q string) (statuses []string, assignee string, searchTerm string) {
	tokens := strings.Split(strings.TrimSpace(q), " ")

	var searchParts []string

	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}

		switch {
		case strings.HasPrefix(token, "status:"):
			raw := strings.TrimPrefix(token, "status:")
			statuses = strings.Split(raw, "|")

		case strings.HasPrefix(token, "assignee:"):
			assignee = strings.TrimPrefix(token, "assignee:")

		default:
			searchParts = append(searchParts, token)
		}
	}

	if len(searchParts) > 0 {
		searchTerm = strings.Join(searchParts, " ")
	}

	return statuses, assignee, searchTerm
}

func (ts *TaskService) UpdateTaskTitle(userID uint, taskID uint, title string) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	task.Title = title

	err = task.Validate()
	if err != nil {
		return models.Task{}, err
	}

	err = ts.db.TaskRepository.Update(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) UpdateTaskDescription(userID uint, taskID uint, description string) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	task.Description = description

	err = task.Validate()
	if err != nil {
		return models.Task{}, err
	}

	err = ts.db.TaskRepository.Update(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) UpdateTaskStatus(userID uint, taskID uint, status string) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	task.Status = status

	err = task.Validate()
	if err != nil {
		return models.Task{}, err
	}

	err = ts.db.TaskRepository.Update(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) UpdateTaskAssignee(userID uint, taskID uint, assigneeID uint) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	if assigneeID != 0 {
		can, _ = ts.rs.CheckRole(assigneeID, task.BoardID, MemberRole)
		if !can {
			return models.Task{}, errors.New("forbidden")
		}
	}

	task.AssigneeID = assigneeID

	err = task.Validate()
	if err != nil {
		return models.Task{}, err
	}

	err = ts.db.TaskRepository.Update(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) UploadFile(userID uint, taskID uint, file []byte, name string) (models.File, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.File{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.File{}, errors.New("forbidden")
	}

	if strings.TrimSpace(name) == "" {
		return models.File{}, errors.New("file name cannot be empty")
	}

	path := uuid.New().String()

	storagePath := fmt.Sprintf("%s/tasks/%d/%s", ts.config.DataDir, task.ID, path)

	err = saveFileToStorage(storagePath, file)
	if err != nil {
		return models.File{}, fmt.Errorf("failed to save file: %w", err)
	}

	fileType := "file"
	// if the file type is an image, set the type to image
	if strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") || strings.HasSuffix(name, ".gif") {
		fileType = "image"
	}

	fileRecord := models.File{
		Name:       name,
		Path:       path,
		TaskID:     task.ID,
		UploadedBy: userID,
		Type:       fileType,
	}

	err = ts.db.FileRepository.Create(&fileRecord)
	if err != nil {
		return models.File{}, fmt.Errorf("failed to create file record: %w", err)
	}

	return fileRecord, nil
}

func saveFileToStorage(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(path, data, os.ModePerm)
}

func (ts *TaskService) GetFile(userID uint, fileID uint) (models.File, []byte, error) {
	file, err := ts.db.FileRepository.GetByID(fileID, repository.WithPreload("Task"))
	if err != nil {
		return models.File{}, nil, err
	}

	can, _ := ts.rs.CheckRole(userID, file.Task.BoardID, MemberRole)
	if !can {
		return models.File{}, nil, errors.New("forbidden")
	}

	filePath := fmt.Sprintf("%s/tasks/%d/%s", ts.config.DataDir, file.Task.ID, file.Path)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return models.File{}, nil, err
	}

	return file, content, nil
}

func (ts *TaskService) DeleteFile(userID uint, fileID uint) (models.File, error) {
	file, err := ts.db.FileRepository.GetByID(fileID, repository.WithPreload("Task"))
	if err != nil {
		return models.File{}, err
	}

	can, _ := ts.rs.CheckRole(userID, file.Task.BoardID, MemberRole)
	if !can {
		return models.File{}, errors.New("forbidden")
	}

	filePath := fmt.Sprintf("%s/tasks/%d/%s", ts.config.DataDir, file.Task.ID, file.Path)

	err = os.Remove(filePath)
	if err != nil {
		return models.File{}, err
	}

	err = ts.db.FileRepository.Delete(file.ID)
	if err != nil {
		return models.File{}, err
	}

	return file, nil
}

func (ts *TaskService) CreateTaskLink(userID uint, srcTaskID uint, dstTaskID uint, linkType string) (models.TaskLinks, error) {
	srcTask, err := ts.db.TaskRepository.GetByID(srcTaskID)
	if err != nil {
		return models.TaskLinks{}, err
	}

	can, _ := ts.rs.CheckRole(userID, srcTask.BoardID, MemberRole)
	if !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}

	dstTask, err := ts.db.TaskRepository.GetByID(dstTaskID)
	if err != nil {
		return models.TaskLinks{}, err
	}

	can, _ = ts.rs.CheckRole(userID, dstTask.BoardID, MemberRole)
	if !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}

	lType, ok := repository.LinkTypeMap[linkType]
	inverseLType, inverseOk := repository.InverseLinkTypeMap[linkType]

	if !ok && !inverseOk {
		return models.TaskLinks{}, errors.New("invalid link type")
	}

	rSrcId := srcTaskID
	rDstId := dstTaskID
	rLType := lType

	if !ok {
		rSrcId = dstTaskID
		rDstId = srcTaskID
		rLType = inverseLType
	}

	link := models.TaskLinks{
		SrcTaskID: rSrcId,
		DstTaskID: rDstId,
		LinkType:  string(rLType),
	}

	err = ts.db.TaskLinkRepository.Create(&link)
	if err != nil {
		return models.TaskLinks{}, err
	}

	return link, nil
}

func (ts *TaskService) DeleteTaskLink(userID uint, linkID uint) (models.TaskLinks, error) {
	link, err := ts.db.TaskLinkRepository.GetByID(linkID)
	if err != nil {
		return models.TaskLinks{}, err
	}

	can, _ := ts.rs.CheckRole(userID, link.SrcTask.BoardID, MemberRole)
	if !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}

	can, _ = ts.rs.CheckRole(userID, link.DstTask.BoardID, MemberRole)
	if !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}

	err = ts.db.TaskLinkRepository.Delete(link.ID)
	if err != nil {
		return models.TaskLinks{}, err
	}

	return link, nil
}

func (ts *TaskService) CreateTaskExternalLink(userID uint, taskID uint, title string, url string) (models.TaskExternalLink, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.TaskExternalLink{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, MemberRole)
	if !can {
		return models.TaskExternalLink{}, errors.New("forbidden")
	}

	link := models.TaskExternalLink{
		URL:    url,
		Title:  title,
		TaskID: taskID,
	}

	err = ts.db.TaskExternalLinkRepository.Create(&link)
	if err != nil {
		return models.TaskExternalLink{}, err
	}

	return link, nil
}

func (ts *TaskService) UpdateTaskExternalLink(userID uint, linkID uint, title string, url string) (models.TaskExternalLink, error) {
	link, err := ts.db.TaskExternalLinkRepository.GetByID(linkID, repository.WithPreload("Task"))
	if err != nil {
		return models.TaskExternalLink{}, err
	}

	can, _ := ts.rs.CheckRole(userID, link.Task.BoardID, MemberRole)
	if !can {
		return models.TaskExternalLink{}, errors.New("forbidden")
	}

	link.URL = url
	link.Title = title

	err = ts.db.TaskExternalLinkRepository.Update(&link)
	if err != nil {
		return models.TaskExternalLink{}, err
	}

	return link, nil
}

func (ts *TaskService) DeleteTaskExternalLink(userID uint, linkID uint) (uint, error) {
	link, err := ts.db.TaskExternalLinkRepository.GetByID(linkID, repository.WithPreload("Task"))
	if err != nil {
		return 0, err
	}

	can, _ := ts.rs.CheckRole(userID, link.Task.BoardID, MemberRole)
	if !can {
		return 0, errors.New("forbidden")
	}

	err = ts.db.TaskExternalLinkRepository.Delete(link.ID)
	if err != nil {
		return 0, err
	}

	return link.TaskID, nil
}
