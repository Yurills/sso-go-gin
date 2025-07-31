package service

import (
	"errors"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/internal/sso/par/dtos"
	"sso-go-gin/pkg/utils/randomutil"
	"time"

	"sso-go-gin/internal/sso/par/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "time"
	// "sso-go-gin/internal/sso/models"
	// "errors"
	// "github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type PARService struct {
	repository *repository.PARRepository
}

func NewPARService(repository *repository.PARRepository) *PARService {
	return &PARService{repository: repository}
}

func (s *PARService) CreateRequestURI(c *gin.Context, req dtos.PARRequest) (*dtos.PARResponse, error) {
	// Validate the request
	client, err := s.repository.GetAuthClientByID(c, req.ClientID)
	if err != nil {
		return nil, errors.New("client not found")
	}
	if !client.Active {
		return nil, errors.New("client is not active")
	}
	//check callback uri
	// if req.RedirectURI != client.SSORedirectCallbackURI {
	// 	return nil, errors.New("redirect URI does not match client configuration")
	// }

	//verify SSO token
	ssoToken, err := s.repository.GetSSOTokenByToken(c, req.SSOToken)
	if err != nil {
		return nil, errors.New("sso token not found")
	}
	if ssoToken.IsExpired() {
		return nil, errors.New("sso token is expired")
	}

	AuthRequestCode := &models.AuthRequestCode{
		ID:                     uuid.New(),
		ClientID:               client.ID,
		ResponseType:           "code",
		State:                  req.State,
		CodeChallenge:          req.CodeChallenge,
		CodeChallengeMethod:    req.CodeChallengeMethod,
		SSORedirectCallbackURI: req.RedirectURI,
		ExpiredDatetime:        time.Now().Add(5 * time.Minute), // Set expiration time
		CreatedDatetime:        time.Now(),
		Nonce:                  nil,
	}
	if err := s.repository.SaveAuthRequest(c, AuthRequestCode); err != nil {
		return nil, errors.New("failed to save auth request")
	}
	// Create SSO request URI

	refVal, err := randomutil.GenerateRandomString(32) // Generate a random request URI
	if err != nil {
		return nil, errors.New("failed to generate request URI")
	}
	requestURI := "urn:ietf:params:oauth:request_uri:" + refVal

	ssoRequestURI := &models.SSORequestURI{
		ID:              uuid.New(),
		ClientID:        client.ID,
		SSOToken:        req.SSOToken,
		User:            ssoToken.User,
		RequestURI:      requestURI,
		AuthRequestID:   AuthRequestCode.ID,
		CreatedDatetime: time.Now(),
		ExpiredDatetime: time.Now().Add(5 * time.Minute), // Set expiration time
	}
	if err := s.repository.SaveSSORequestURI(c, ssoRequestURI); err != nil {
		return nil, errors.New("failed to save SSO request URI")
	}

	return &dtos.PARResponse{
		RequestURI: ssoRequestURI.RequestURI,
	}, nil

}

func (s *PARService) GenerateRefreshToken(c *gin.Context, sessionID string, req dtos.PARRequest) (*dtos.PARSessionResponse, error) {
	// Validate the request
	authClient, err := s.repository.GetAuthClientByID(c, req.ClientID)
	if err != nil {
		return nil, errors.New("client not found")
	}
	if !authClient.Active {
		return nil, errors.New("client is not active")
	}

	//verify session
	user, err := s.repository.GetUserInfoBySessionID(c, sessionID)
	if err != nil {
		return nil, errors.New("failed to get user info")
	}

	refreshtoken, err := randomutil.GenerateRandomString(32)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	refreshToken := &models.RefreshToken{
		ID:              uuid.New(),
		RefreshToken:    refreshtoken,
		ClientID:        authClient.ID,
		User:            user.Username,
		Email:           user.Email,
		ExpiredDatetime: time.Now().Add(1 * time.Hour), // Set expiration time for the refresh token
		CreatedDatetime: time.Now(),
	}
	if err := s.repository.SaveRefreshToken(c, refreshToken); err != nil {
		return nil, errors.New("failed to save refresh token")
	}

	response := &dtos.PARSessionResponse{
		RefreshToken: refreshtoken,
		RedirectURI:  authClient.AuthRedirectCallbackURI,
		State:        req.State,
	}
	return response, nil

}
