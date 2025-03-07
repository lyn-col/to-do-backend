package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TaskController struct holds DB and logger
type TaskController struct {
	DB     *gorm.DB
	Logger *log.Logger
}

// NewTaskController initializes TaskController with dependencies
func NewTaskController(db *gorm.DB, logger *log.Logger) *TaskController {
	return &TaskController{DB: db, Logger: logger}
}

// GetTasks retrieves all tasks
func (tc *TaskController) GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tasks": []string{"Task 1", "Task 2"}})
	// Implement database query for tasks
}

// CreateTask creates a new task
func (tc *TaskController) CreateTask(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
	// Implement logic
}

// UpdateTask updates an existing task
func (tc *TaskController) UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	// Implement logic
}

// DeleteTask removes a task
func (tc *TaskController) DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	// Implement logic
}