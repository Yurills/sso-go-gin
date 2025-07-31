package register

import (
	"context"
	"errors"
	"sso-go-gin/internal/auth/models"
	"sso-go-gin/pkg/utils/hashutil"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repository *repository
}

func NewService(repo *repository) *Service {
	return &Service{repo}
}

func (s *Service) Register(c context.Context, req RegisterRequest) (*models.User, error) {

	// check if the username already exists
	existingUser, err := s.repository.GetUserInfo(c, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := hashutil.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create a new user model
	user := &models.User{
		ID:                 uuid.New(), // Assuming GenerateUUID generates a new UUID
		Username:           req.Username,
		Password:           hashedPassword,
		Email:              req.Email,
		TwoFAEnabled:       false, // Default value
		ForceResetPassword: false, // Default value

	}

	// Save the user to the database
	createdUser, err := s.repository.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
