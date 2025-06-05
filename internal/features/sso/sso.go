package sso

import (
	"errors"
	"sso-go-gin/internal/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	repository *repository
}

func NewService(repo *repository) *Service {
	return &Service{repo}
}

func (s *Service) Login(ctx *gin.Context, req LoginRequest) (*LoginResponse, error) {
	//verify require parameters
	if req.RID == "" || req.Username == "" || req.Password == "" {
		return nil, errors.New("missing required parameters")
	}

	//check if RID is provided and not expired
	authReq, err := s.repository.GetAuthRequestByID(ctx, req.RID)
	if err != nil || authReq.IsExpired() {
		return nil, errors.New("invalid or expired auth request")
	}

	//check if credentials are valid
	user, err := s.repository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	//generate auth code and insert into database
	authCode := utils.GenerateRandomString(32) // Generate a random auth code
	authCodeRecord := &AuthCode{
		ID:              uuid.New(),
		Code:            authCode,
		RID:             uuid.MustParse(req.RID),
		Type:            "code",
		ExpiredDatetime: time.Now().Add(5 * time.Minute),
		CreatedDatetime: time.Now(),
	}
	if err := s.repository.SaveAuthCode(ctx, authCodeRecord); err != nil {
		return nil, errors.New("failed to save auth code")
	}

	loginResponse := &LoginResponse{
		AuthCode:    authCodeRecord.Code,
		RedirectURI: authReq.AuthRedirectCallbackURI,
		State:       authReq.State,
	}

	//create cookie
	return loginResponse, nil

}
