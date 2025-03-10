package server

import (
	"log"

	"to-do-backend/controllers"
	"to-do-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes initializes API routes with controllers
func SetupRoutes(r *gin.Engine, db *gorm.DB, logger *log.Logger) {
	// Initialize controllers
	authService := services.NewAuthService(db)
	taskService := services.NewTaskService(db)

	// Initialize controllers with services
	authController := controllers.NewAuthController(db, authService, logger)
	taskController := controllers.NewTaskController(db, taskService, logger)

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.RegisterUser)
		auth.POST("/login", authController.LoginUser)
	}

	// Task routes
	tasks := r.Group("/tasks")
	{
		tasks.GET("", taskController.GetTasks)
		tasks.POST("", taskController.CreateTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}
}
