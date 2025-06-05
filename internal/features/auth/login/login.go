package login

import (
	"context"
	"errors"
	"sso-go-gin/internal/models"
	"sso-go-gin/internal/pkg/utils"
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

	if utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("wrong password")
	}

	return user, nil

}
