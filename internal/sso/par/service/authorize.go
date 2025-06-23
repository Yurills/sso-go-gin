package service

import (
	"errors"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/internal/sso/par/dtos"
	"sso-go-gin/pkg/utils/randomutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *PARService) Authorize(c *gin.Context, req dtos.PARRequestAuthorize) (*dtos.PARResponseAuthorize, error) {
	// Validate the request
	ssoRequestURI, err := s.repository.GetSSOTokenByRequestURI(c, req.RequestURI)
	if err != nil {
		return nil, errors.New("SSO request URI not found")
	}
	if ssoRequestURI.IsExpired() {
		return nil, errors.New("SSO request URI is expired")
	}

	// Create the response
	authRequest, err := s.repository.GetAuthRequestByID(c, ssoRequestURI.AuthRequestID.String())
	if err != nil {
		return nil, errors.New("authorization request not found")
	}

	authCode, err := randomutil.GenerateRandomString(32) // Generate a random auth code
	if err != nil {
		return nil, errors.New("failed to generate auth code")
	}

	authCodeRecord := &models.AuthCode{
		ID:              uuid.New(),
		Code:            authCode,
		RID:             authRequest.ID,
		Type:            "code",
		ExpiredDatetime: time.Now().Add(5 * time.Minute),
		CreatedDatetime: time.Now(),
		Username:        ssoRequestURI.User,
	}
	if err := s.repository.SaveAuthCode(c, authCodeRecord); err != nil {
		return nil, errors.New("failed to save authorization code")
	}

	response := &dtos.PARResponseAuthorize{
		Code:        authCode,
		RedirectURI: authRequest.SSORedirectCallbackURI,
	}
	return response, nil
}
