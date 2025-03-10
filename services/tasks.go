package services

import (
	"errors"
	"to-do-backend/models"
	"to-do-backend/dto"
	"gorm.io/gorm"
)

// TaskService defines task-related operations
type TaskService interface {
	GetTasks(userID uint) ([]models.Task, error)
	CreateTask(task *dto.Task) error
	UpdateTask(id uint, updatedTask *models.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	db *gorm.DB
}

// NewTaskService creates an instance of TaskService
func NewTaskService(db *gorm.DB) TaskService {
	return &taskService{db}
}

func (s *taskService) GetTasks(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (s *taskService) CreateTask(task *dto.Task) error {
	return s.db.Create(task).Error
}

func (s *taskService) UpdateTask(id uint, updatedTask *models.Task) error {
	var task models.Task
	if err := s.db.First(&task, id).Error; err != nil {
		return errors.New("task not found")
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Completed = updatedTask.Completed

	return s.db.Save(&task).Error
}

func (s *taskService) DeleteTask(id uint) error {
	if err := s.db.Delete(&models.Task{}, id).Error; err != nil {
		return errors.New("task not found")
	}
	return nil
}
