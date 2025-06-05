package sso

import (
	"context"
	"errors"
	"sso-go-gin/internal/pkg/utils"
	"time"
)

type Service struct {
	repository *repository
}

func NewService(repo *repository) *Service {
	return &Service{repo}
}


func (s *Service) Login(ctx context.Context, username, password string) (*User, error) {
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}


// authorization code flow
func (s *Service) Authorize(c context.Context, rid string) (*AuthRequestCode, error) {
	req, err := s.repository.GetAuthRequestByID(c, rid)
	if err != nil || time.Now().After(req.ExpiresAt) {
		return nil, errors.New("authorziation request expired or not found")
	}
	return req, nil
}

// send auth code
func (s *Service) IssueAuthCode(c context.Context, userId uint, rid string) (*AuthCode, string, error) {
	req, err := s.repository.GetAuthRequestByID(c, rid)
	if err != nil {
		return nil, "", err
	}

	code := &AuthCode{
		Code:      utils.GenerateRandomString(32),
		ClientID:  string(rune(userId)),
		Type:      "code",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	if err := s.repository.SaveAuthCode(c, *code); err != nil {
		return nil, "", err
	}

	redirectURI := req.RedirectURI + "?code=" + code.Code + "&state=" + req.State
	return code, redirectURI, nil
}
