package task

import (
	"encoding/base64"
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/models"
	"server/services/board"
	"server/services/eventBus"
	"server/services/role"
	"server/services/task"
	"strings"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	db  *repository.Database
	ts  *task.TaskService
	rs  *role.RoleService
	bs  *board.BoardService
	hs  *helpers.HelperService
	te  *eventBus.EventBus[models.Task]
	fe  *eventBus.EventBus[models.File]
	le  *eventBus.EventBus[models.TaskLinks]
	lee *eventBus.EventBus[models.TaskExternalLink]
}

func NewTaskController(db *repository.Database, ts *task.TaskService, rs *role.RoleService, bs *board.BoardService, hs *helpers.HelperService, te *eventBus.EventBus[models.Task], fe *eventBus.EventBus[models.File], le *eventBus.EventBus[models.TaskLinks], lee *eventBus.EventBus[models.TaskExternalLink]) *TaskController {
	return &TaskController{db: db, ts: ts, rs: rs, bs: bs, hs: hs, te: te, fe: fe, le: le, lee: lee}
}

// @Summary Create a task
// @Description Create a task
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateTaskRequest true "Create task request"
// @Success 200 {object} CreateTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/create [post]
func (tc *TaskController) CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.CreateTask(user.ID, task.CreateTaskRequest{
		ParentTaskID: request.ParentTaskID,
		BoardID:      request.BoardID,
		Title:        request.Title,
		Description:  request.Description,
		SwimlaneID:   request.SwimlaneID,
		ColumnID:     request.ColumnID,
		Status:       request.Status,
		AssigneeID:   request.AssigneeID,
		DueDate:      request.DueDate,
	})
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	tc.te.Publish("task.created", models.Task{}, task, user)

	c.JSON(http.StatusOK, CreateTaskResponse{Task: task})
}

// @Summary Delete a task
// @Description Delete a task
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteTaskRequest true "Delete task request"
// @Success 200 {object} DeleteTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/delete [post]
func (tc *TaskController) DeleteTask(c *gin.Context) {
	var request DeleteTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.DeleteTaskRequest(user.ID, request.ID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	tc.te.Publish("task.deleted", task, models.Task{}, user)

	c.JSON(http.StatusOK, DeleteTaskResponse{Task: task})
}

// @Summary Get a task
// @Description Get a task
// @Tags tasks
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param id path uint true "Task ID"
// @Success 200 {object} GetTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/get/{id} [get]
func (tc *TaskController) GetTask(c *gin.Context) {
	var request GetTaskRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.GetTask(user.ID, request.ID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, GetTaskResponse{Task: task})
}

// @Summary Move a task
// @Description Move a task to a different column and swimlane, and update its position
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body MoveTaskRequest true "Move task request"
// @Success 200 {object} MoveTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/move [post]
func (tc *TaskController) MoveTask(c *gin.Context) {
	var request MoveTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, oldTask, err := tc.ts.MoveTask(user.ID, task.MoveTaskRequest{
		TaskID:     request.TaskID,
		ColumnID:   request.ColumnID,
		SwimlaneID: request.SwimlaneID,
		Position:   request.Position,
	})
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	tc.te.Publish("task.moved", oldTask, task, user)

	c.JSON(http.StatusOK, MoveTaskResponse{Task: task})
}

// @Summary Get tasks with a query
// @Description Get tasks with a query
// @Tags tasks
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param query path string true "Query"
// @Param board_id path uint true "Board ID"
// @Success 200 {object} GetTaskQueryResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/get-query/{board_id}/{query} [get]
func (tc *TaskController) GetTaskQuery(c *gin.Context) {
	var request GetTaskQueryRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tasks, err := tc.ts.GetTasksWithQuery(user.ID, request.BoardID, request.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetTaskQueryResponse{Tasks: tasks})
}

type UploadFileRequest struct {
	TaskID uint   `json:"task_id"`
	File   []byte `json:"file"`
	Name   string `json:"name"`
}

type UploadFileResponse struct {
	File models.File `json:"file"`
}

// @Summary Upload a file
// @Description Upload a file to a task
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UploadFileRequest true "Upload file request"
// @Success 200 {object} UploadFileResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/upload [post]
func (tc *TaskController) UploadFile(c *gin.Context) {
	var request UploadFileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := tc.ts.UploadFile(user.ID, request.TaskID, request.File, request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.fe.Publish("file.created", models.File{}, file, user)

	c.JSON(http.StatusOK, UploadFileResponse{File: file})
}

type GetImageRequest struct {
	FileID uint `uri:"file_id"`
}

type GetImageResponse struct {
	File    models.File `json:"file"`
	Content string      `json:"content"`
}

// @Summary Get an image
// @Description Get an image
// @Tags tasks
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param file_id path uint true "File ID"
// @Success 200 {object} GetImageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/get-image/{file_id} [get]
func (tc *TaskController) GetImage(c *gin.Context) {
	var request GetImageRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, content, err := tc.ts.GetFileRequest(user.ID, request.FileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if file.Type != "image" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
		return
	}

	contentString := base64.StdEncoding.EncodeToString(content)

	if strings.HasSuffix(file.Name, ".png") {
		contentString = "data:image/png;base64," + contentString
	} else if strings.HasSuffix(file.Name, ".jpg") || strings.HasSuffix(file.Name, ".jpeg") {
		contentString = "data:image/jpeg;base64," + contentString
	} else if strings.HasSuffix(file.Name, ".gif") {
		contentString = "data:image/gif;base64," + contentString
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
		return
	}

	response := GetImageResponse{
		File:    file,
		Content: contentString,
	}

	c.JSON(http.StatusOK, response)
}

type DownloadFileRequest struct {
	FileID uint `uri:"file_id"`
}

// todo serve file as binary
type DownloadFileResponse struct {
	File    models.File `json:"file"`
	Content string      `json:"content"`
}

// @Summary Download a file
// @Description Download a file
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param file_id path uint true "File ID"
// @Success 200 {object} DownloadFileResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/download/{file_id} [get]
func (tc *TaskController) DownloadFile(c *gin.Context) {
	var request DownloadFileRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, content, err := tc.ts.GetFileRequest(user.ID, request.FileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DownloadFileResponse{File: file, Content: base64.StdEncoding.EncodeToString(content)})
}

type DeleteFileRequest struct {
	FileID uint `json:"file_id"`
}

type DeleteFileResponse struct {
	File models.File `json:"file"`
}

// @Summary Delete a file
// @Description Delete a file
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteFileRequest true "Delete file request"
// @Success 200 {object} DeleteFileResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/delete-file [post]
func (tc *TaskController) DeleteFile(c *gin.Context) {
	var request DeleteFileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := tc.ts.DeleteFileRequest(user.ID, request.FileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.fe.Publish("file.deleted", file, models.File{}, user)

	c.JSON(http.StatusOK, DeleteFileResponse{File: file})
}

type CreateTaskLinkRequest struct {
	SrcTaskID uint   `json:"src_task_id"`
	DstTaskID uint   `json:"dst_task_id"`
	LinkType  string `json:"link_type" binding:"required,oneof=depends_on blocks fixes depended_on_by blocked_by fixed_by"`
}

type CreateTaskLinkResponse struct {
	Link models.TaskLinks `json:"link"`
}

// @Summary Create a task link
// @Description Create a task link
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateTaskLinkRequest true "Create task link request"
// @Success 200 {object} CreateTaskLinkResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/create-link [post]
func (tc *TaskController) CreateTaskLink(c *gin.Context) {
	var request CreateTaskLinkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if request.SrcTaskID == request.DstTaskID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Source and destination task IDs cannot be the same"})
		return
	}

	link, err := tc.ts.CreateTaskLink(user.ID, request.SrcTaskID, request.DstTaskID, request.LinkType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.le.Publish("link.created", models.TaskLinks{}, link, user)

	c.JSON(http.StatusOK, CreateTaskLinkResponse{Link: link})
}

type DeleteTaskLinkRequest struct {
	LinkID uint `json:"link_id"`
}

type DeleteTaskLinkResponse struct {
	Link models.TaskLinks `json:"link"`
}

// @Summary Delete a task link
// @Description Delete a task link
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteTaskLinkRequest true "Delete task link request"
// @Success 200 {object} DeleteTaskLinkResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/delete-link [post]
func (tc *TaskController) DeleteTaskLink(c *gin.Context) {
	var request DeleteTaskLinkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	link, err := tc.ts.DeleteTaskLinkRequest(user.ID, request.LinkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.le.Publish("link.deleted", link, models.TaskLinks{}, user)

	c.JSON(http.StatusOK, DeleteTaskLinkResponse{Link: link})
}

type QueryAllBoardsRequest struct {
	Query string `json:"query"`
}

type QueryAllBoardsResponse struct {
	Tasks []models.Task `json:"tasks"`
}

// @Summary Simple query all boards
// @Description Simple query all boards
// @Tags tasks
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param request body QueryAllBoardsRequest true "Query all boards request"
// @Success 200 {object} QueryAllBoardsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/query-all-boards [post]
func (tc *TaskController) QueryAllBoards(c *gin.Context) {
	var request QueryAllBoardsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	boards, err := tc.bs.ListBoards(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tasks := []models.Task{}
	for _, board := range boards {
		t, err := tc.ts.GetTasksWithQuery(user.ID, board.ID, request.Query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, t...)
	}

	c.JSON(http.StatusOK, QueryAllBoardsResponse{Tasks: tasks})
}

type CreateTaskExternalLinkRequest struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	TaskID uint   `json:"task_id"`
}

type CreateTaskExternalLinkResponse struct {
	Link models.TaskExternalLink `json:"link"`
}

// @Summary Create a task external link
// @Description Create a task external link
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateTaskExternalLinkRequest true "Create task external link request"
// @Success 200 {object} CreateTaskExternalLinkResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/create-external-link [post]
func (tc *TaskController) CreateTaskExternalLink(c *gin.Context) {
	var request CreateTaskExternalLinkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	link, err := tc.ts.CreateTaskExternalLink(user.ID, request.TaskID, request.Title, request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.lee.Publish("externallink.created", models.TaskExternalLink{}, link, user)

	c.JSON(http.StatusOK, CreateTaskExternalLinkResponse{Link: link})
}

type UpdateTaskExternalLinkRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type UpdateTaskExternalLinkResponse struct {
	Link models.TaskExternalLink `json:"link"`
}

// @Summary Update a task external link
// @Description Update a task external link
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateTaskExternalLinkRequest true "Update task external link request"
// @Success 200 {object} UpdateTaskExternalLinkResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/update-external-link [post]
func (tc *TaskController) UpdateTaskExternalLink(c *gin.Context) {
	var request UpdateTaskExternalLinkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	link, oldLink, err := tc.ts.UpdateTaskExternalLink(user.ID, request.ID, request.Title, request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.lee.Publish("externallink.updated", oldLink, link, user)

	c.JSON(http.StatusOK, UpdateTaskExternalLinkResponse{Link: link})
}

type DeleteTaskExternalLinkRequest struct {
	ID uint `json:"id"`
}

type DeleteTaskExternalLinkResponse struct {
	TaskID  uint   `json:"task_id"`
	Message string `json:"message"`
}

// @Summary Delete a task external link
// @Description Delete a task external link
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteTaskExternalLinkRequest true "Delete task external link request"
// @Success 200 {object} DeleteTaskExternalLinkResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/delete-external-link [post]
func (tc *TaskController) DeleteTaskExternalLink(c *gin.Context) {
	var request DeleteTaskExternalLinkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	link, err := tc.ts.DeleteTaskExternalLinkRequest(user.ID, request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.lee.Publish("externallink.deleted", link, models.TaskExternalLink{}, user)

	c.JSON(http.StatusOK, DeleteTaskExternalLinkResponse{TaskID: link.TaskID, Message: "deleted"})
}
