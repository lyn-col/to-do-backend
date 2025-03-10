package controllers

import (
	"log"
	"net/http"
	"to-do-backend/dto"
	"to-do-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthController struct with dependencies
type AuthController struct {
	DB          *gorm.DB
	AuthService services.AuthService
	Logger      *log.Logger
}

// NewAuthController initializes AuthController with AuthService and DB
func NewAuthController(db *gorm.DB, authService services.AuthService, logger *log.Logger) *AuthController {
	return &AuthController{DB: db, AuthService: authService, Logger: logger}
}

// RegisterUser registers a new user
func (ac *AuthController) RegisterUser(c *gin.Context) {
	ac.Logger.Println("Registering user...")

	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.AuthService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser logs in a user and returns a JWT token
func (ac *AuthController) LoginUser(c *gin.Context) {
	ac.Logger.Println("Logging in user...")

	var loginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.AuthService.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a placeholder JWT token (Replace with actual JWT generation logic)
	token := "jwt_token_placeholder"

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
}
