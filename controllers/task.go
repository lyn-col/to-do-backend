package controllers

import (
	"log"
	"net/http"
	"strconv"
	"to-do-backend/models"
	"to-do-backend/services"
	"to-do-backend/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TaskController struct with dependencies
type TaskController struct {
	DB          *gorm.DB
	TaskService services.TaskService
	Logger      *log.Logger
}

// NewTaskController initializes TaskController with TaskService and DB
func NewTaskController(db *gorm.DB, taskService services.TaskService, logger *log.Logger) *TaskController {
	return &TaskController{DB: db, TaskService: taskService, Logger: logger}
}

// GetTasks retrieves tasks for a user
func (tc *TaskController) GetTasks(c *gin.Context) {
	tc.Logger.Println("Fetching tasks...")

	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := tc.TaskService.GetTasks(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CreateTask creates a new task
func (tc *TaskController) CreateTask(c *gin.Context) {
	tc.Logger.Println("Creating a new task...")

	var task dto.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TaskService.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task
func (tc *TaskController) UpdateTask(c *gin.Context) {
	tc.Logger.Println("Updating task...")

	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TaskService.UpdateTask(uint(taskID), &updatedTask); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask deletes a task by ID
func (tc *TaskController) DeleteTask(c *gin.Context) {
	tc.Logger.Println("Deleting task...")

	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := tc.TaskService.DeleteTask(uint(taskID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
