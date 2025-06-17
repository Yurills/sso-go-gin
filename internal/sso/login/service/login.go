package service

import (
	"errors"
	"net/http"
	"time"

	"sso-go-gin/internal/sso/login/dtos"
	"sso-go-gin/internal/sso/login/repository"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/pkg/utils/randomutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	user, err := s.repository.GetUserInfo(ctx, req.Username)
	if err != nil || user == nil {
		return nil, errors.New("invalid username or password")
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("wrong password")
	}

	//generate auth code and insert into database
	authCode, err := randomutil.GenerateRandomString(32) // Generate a random auth code
	if err != nil {
		return nil, errors.New("failed to generate auth code")
	}

	authCodeRecord := &models.AuthCode{
		ID:              uuid.New(),
		Code:            authCode,
		RID:             uuid.MustParse(req.RID),
		Type:            "code",
		ExpiredDatetime: time.Now().Add(5 * time.Minute),
		CreatedDatetime: time.Now(),
		Username:        req.Username,
	}
	if err := s.repository.SaveAuthCode(ctx, authCodeRecord); err != nil {
		return nil, errors.New("failed to save auth code")
	}

	//create cookie on browser
	session_id := uuid.New().String()
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:  "sso_session",
		Value: session_id,
		Path:  "/",
		// HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour), // Set cookie expiration
	})

	loginResponse := &dtos.LoginResponse{
		AuthCode:    authCodeRecord.Code,
		RedirectURI: authReq.AuthRedirectCallbackURI,
		State:       authReq.State,
		Nonce:       authReq.Nonce,
	}
	return loginResponse, nil

}
