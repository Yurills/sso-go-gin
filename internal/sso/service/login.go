package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sso-go-gin/pkg/utils/randomutil"
	"sso-go-gin/internal/sso/dtos"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/internal/sso/repository"
)

type LoginService struct {
	repository *repository.LoginRepository
}

func NewLoginService(repo *repository.LoginRepository) *LoginService {
	return &LoginService{repo}
}

func (s *LoginService) Login(ctx *gin.Context, req dtos.LoginRequest) (*dtos.LoginResponse, error) {
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
	

	
	//generate auth code and insert into database
	authCode := randomutil.GenerateRandomString(32) // Generate a random auth code
	authCodeRecord := &models.AuthCode{
		ID:              uuid.New(),
		Code:            authCode,
		RID:             uuid.MustParse(req.RID),
		Type:            "code",
		ExpiredDatetime: time.Now().Add(5 * time.Minute),
		CreatedDatetime: time.Now(),
	}
	if err := s.repository.SaveAuthCode(ctx, *authCodeRecord); err != nil {
		return nil, errors.New("failed to save auth code")
	}

	loginResponse := &dtos.LoginResponse{
		AuthCode:    authCodeRecord.Code,
		RedirectURI: authReq.AuthRedirectCallbackURI,
		State:       authReq.State,
	}

	//create cookie
	return loginResponse, nil

}
