package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"sso-go-gin/internal/sso/login/dtos"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/pkg/utils/randomutil"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// in the case flow get interrupt by 2FA or password reset, we continue after getting sendback to here.
func (s *LoginService) ContinueLogin(ctx *gin.Context) (*dtos.LoginResponse, error) {
	flow_session := sessions.Default(ctx)
	userID := flow_session.Get("temp_user_id")
	if userID == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	// Continue the login process (e.g. verify 2FA code)
	println("Continuing login for user ID: " + userID.(string))
	println("Session state: " + flow_session.Get("login_state").(string))

	oauth := flow_session.Get("oauth_pending")
	if oauth == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no oauth pending"})
		return nil, errors.New("no oauth pending")
	}
	println(oauth.(string))
	var oauthMap map[string]string
	if err := json.Unmarshal([]byte(oauth.(string)), &oauthMap); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse oauth pending data"})
		return nil, err
	}

	//generate auth code
	authCode, err := randomutil.GenerateRandomString(32) // Generate a random auth code
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate auth code"})
		return nil, err
	}
	authCodeRecord := &models.AuthCode{
		ID:              uuid.New(),
		Code:            authCode,
		RID:             uuid.MustParse(oauthMap["rid"]),
		Type:            "code",
		ExpiredDatetime: time.Now().Add(24 * time.Hour), // Set expiration time
		CreatedDatetime: time.Now(),
		Username:        flow_session.Get("temp_username").(string), // Use temporary username
	}
	if err := s.repository.SaveAuthCode(ctx, authCodeRecord); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save auth code"})
		return nil, err
	}

	//save session
	
	sessionID := uuid.New().String()
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	session := &models.Session{
		ID:              uuid.MustParse(sessionID),
		UserID:          uuid.MustParse(userID.(string)),
		CreatedDatetime: time.Now(),
		ExpiredDatetime: time.Now().Add(24 * time.Hour),
	}
	if err := s.repository.SaveSession(ctx, session); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return nil, err
	}

	LoginResponse := &dtos.LoginResponse{
		AuthCode:    authCodeRecord.Code,
		RedirectURI: oauthMap["redirect_uri"],
		State:       oauthMap["state"],
	}

	// Clear session data
	flow_session.Set("user_id", userID.(string))
	flow_session.Delete("temp_user_id")          
	flow_session.Delete("temp_username")         
	flow_session.Delete("login_state")           
	flow_session.Delete("oauth_pending")
	// ctx.JSON(http.StatusOK, LoginResponse)
	return LoginResponse, nil

}
