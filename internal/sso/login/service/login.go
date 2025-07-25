package service

import (
	"errors"
	"net/http"

	"time"

	"sso-go-gin/internal/sso/login/dtos"
	"sso-go-gin/internal/sso/login/repository"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/pkg/utils/randomutil"

	"github.com/gin-contrib/sessions"
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
	flow_session := sessions.Default(ctx)
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

	//check csrf token
	csrfCookie, err := ctx.Request.Cookie("csrf_token")
	if err != nil {
		return nil, errors.New("missing CSRF token")
	}

	csrfHeader := ctx.GetHeader("X-csrf_token")
	if csrfHeader != csrfCookie.Value {
		return nil, errors.New("invalid CSRF token")
	}

	flow_session.Set("temp_user_id", user.ID.String()) // Temporarily store user ID in session
	flow_session.Set("temp_username", user.Username) // Temporarily store username in session
	//check if user need 2FA or password-reset (break the login flow if needed)
	if user.IsTwoFactorEnabled() {
		// If 2FA is enabled, redirect to 2FA verification page
		flow_session.Set("login_state", "2fa_required")
		flow_session.Save()
		print("session state set:" + flow_session.Get("login_state").(string))
		return nil, errors.New("2FA required")
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
		ExpiredDatetime: time.Now().Add(24 * time.Hour), // Set expiration time
		CreatedDatetime: time.Now(),
		Username:        req.Username,
	}
	if err := s.repository.SaveAuthCode(ctx, authCodeRecord); err != nil {
		return nil, errors.New("failed to save auth code")
	}

	//create cookie on browser
	sessionID := uuid.New().String()
	session := &models.Session{
		ID:              uuid.MustParse(sessionID),
		UserID:          user.ID,
		CreatedDatetime: time.Now(),
		ExpiredDatetime: time.Now().Add(24 * time.Hour),
	}
	if err := s.repository.SaveSession(ctx, session); err != nil {
		return nil, errors.New("failed to save session")
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	print("cookie saved with session ID:", sessionID)

	loginResponse := &dtos.LoginResponse{
		AuthCode:    authCodeRecord.Code,
		RedirectURI: authReq.AuthRedirectCallbackURI,
		State:       authReq.State,
		Nonce:       authReq.Nonce,
	}

	//all clear, set session state to logged in
	flow_session.Set("user_id", user.ID.String())
	flow_session.Delete("temp_user_id") // Clear temporary user ID
	flow_session.Delete("login_state")  // Clear login state
	flow_session.Delete("temp_username") // Clear temporary username
	flow_session.Delete("oauth_pending") // Clear any pending OAuth state
	flow_session.Save()

	return loginResponse, nil

}
