package server

import (
	"log"

	"to-do-backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes initializes API routes with controllers
func SetupRoutes(r *gin.Engine, db *gorm.DB, logger *log.Logger) {
	// Initialize controllers
	authController := controllers.NewAuthController(db, logger)
	taskController := controllers.NewTaskController(db, logger)

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.RegisterUser)
		auth.POST("/login", authController.LoginUser)
	}

	// Task routes
	tasks := r.Group("/tasks")
	{
		tasks.GET("/", taskController.GetTasks)
		tasks.POST("/", taskController.CreateTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}
}