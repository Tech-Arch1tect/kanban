package task

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"server/database/repository"
	"server/models"
	"server/services/role"

	"github.com/google/uuid"
)

func (ts *TaskService) UploadFile(userID uint, taskID uint, file []byte, name string) (models.File, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.File{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.File{}, errors.New("forbidden")
	}
	if strings.TrimSpace(name) == "" {
		return models.File{}, errors.New("file name cannot be empty")
	}

	path := uuid.New().String()
	storagePath := fmt.Sprintf("%s/tasks/%d/%s", ts.config.DataDir, task.ID, path)
	if err = saveFileToStorage(storagePath, file); err != nil {
		return models.File{}, fmt.Errorf("failed to save file: %w", err)
	}

	fileType := "file"
	if strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".jpg") ||
		strings.HasSuffix(name, ".jpeg") || strings.HasSuffix(name, ".gif") {
		fileType = "image"
	}

	fileRecord := models.File{
		Name:       name,
		Path:       path,
		TaskID:     task.ID,
		UploadedBy: userID,
		Type:       fileType,
	}
	if err = ts.db.FileRepository.Create(&fileRecord); err != nil {
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
	file, err := ts.db.FileRepository.GetByID(fileID, repository.WithPreload("Task"), repository.WithPreload("UploadedByUser"), repository.WithPreload("Task.Board"))
	if err != nil {
		return models.File{}, nil, err
	}

	can, _ := ts.rs.CheckRole(userID, file.Task.BoardID, role.MemberRole)
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

	can, _ := ts.rs.CheckRole(userID, file.Task.BoardID, role.MemberRole)
	if !can {
		return models.File{}, errors.New("forbidden")
	}

	filePath := fmt.Sprintf("%s/tasks/%d/%s", ts.config.DataDir, file.Task.ID, file.Path)
	if err = os.Remove(filePath); err != nil {
		return models.File{}, err
	}
	if err = ts.db.FileRepository.Delete(file.ID); err != nil {
		return models.File{}, err
	}
	return file, nil
}
