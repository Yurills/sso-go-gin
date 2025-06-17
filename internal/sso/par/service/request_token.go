package service

import (
	"errors"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/internal/sso/par/dtos"
	"sso-go-gin/pkg/utils/randomutil"
	"sso-go-gin/pkg/utils/tokenutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *PARService) GetRequestToken(c *gin.Context, req *dtos.PARRequestTokenRequest) (*dtos.PARRequestTokenResponse, error) {
	// authRequest, err := s.repository.GetAuthRequestByClientID(c, req.ClientID)
	// if err != nil {
	// 	return nil, err
	// }

	// if authRequest == nil {
	// 	return nil, errors.New("authorization request not found for client ID: " + req.ClientID)
	// }

	// if authRequest.IsExpired() {
	// 	return nil, errors.New("authorization request is expired for client ID: " + req.ClientID)
	// }

	authClient, err := s.repository.GetAuthClientByID(c, req.ClientID)
	if err != nil {
		return nil, errors.New("failed to get auth client: " + err.Error())
	}
	if !authClient.Active {
		return nil, errors.New("auth client is not active for client ID: " + req.ClientID)
	}

	token, err := randomutil.GenerateRandomString(32)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("authorization header is missing or invalid")
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := tokenutil.ParseAndValidateToken(tokenStr)
	if err != nil {
		return nil, errors.New("failed to parse or validate token: " + err.Error())
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return nil, errors.New("subject (sub) claim is missing or invalid in the token")
	}

	sso_token := &models.SSOToken{
		ID:              uuid.New(),
		Token:           token, // Generate a random token
		ClientID:        uuid.MustParse(req.ClientID),
		Source:          req.Source,
		Destination:     req.Destination,
		ExpiredDatetime: time.Now().Add(60 * time.Second), // Set expiration time to 60 seconds
		User:            sub,
	}

	if err := s.repository.SaveSSOToken(c, sso_token); err != nil {
		return nil, errors.New("failed to save SSO token: " + err.Error())
	}
	response := &dtos.PARRequestTokenResponse{
		Token: token,
	}
	return response, nil

}
