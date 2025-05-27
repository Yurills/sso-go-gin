package login

import (
	"context"
	"errors"
	"sso-go-gin/internal/features/models"
)

type Service struct {
	repository *repository
	//create token maker
}

func NewService(repo *repository) *Service {
	return &Service{repo}
}

func (s *Service) Login(c context.Context, req LoginRequest) (*models.User, error) {

	//check if user matches
	user, err := s.repository.GetUserInfo(c, req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.Password != req.Password {
		return nil, errors.New("wrong Password")
	}

	return user, nil

}
