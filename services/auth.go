package services

import (
	"errors"
	"fmt"
	"to-do-backend/dto"
	"to-do-backend/models"

	"gorm.io/gorm"
)

// AuthService defines authentication methods
type AuthService interface {
	RegisterUser(user *dto.User) error
	LoginUser(email, password string) (*models.User, error)
}

type authService struct {
	db *gorm.DB
}

// NewAuthService creates an instance of AuthService
func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db}
}

func (s *authService) RegisterUser(user *dto.User) error {
	u := models.User{Email: user.Email, Username: user.Username, Password: user.Password}
	return s.db.Create(&u).Error
}

func (s *authService) LoginUser(email, password string) (*models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	if password != user.Password {
		fmt.Println(password)
		fmt.Println(user.Password)
		return nil, errors.New("invalid cred")
	}
	return &user, nil
}
