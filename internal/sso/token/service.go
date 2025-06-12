package token

import (
	"errors"
	"sso-go-gin/pkg/utils/hashutil"
	"sso-go-gin/pkg/utils/randomutil"
	"sso-go-gin/pkg/utils/tokenutil"

	"github.com/gin-gonic/gin"
)

type TokenService struct {
	repository *TokenRepository
}

func NewTokenService(repository *TokenRepository) *TokenService {
	return &TokenService{repository}
}

func (s *TokenService) GenerateToken(ctx *gin.Context, req TokenRequest) (*TokenResponse, error) {

	//verify grant type
	if req.GrantType != "authorization_code" {
		return nil, errors.New("invalid grant type")
	}

	client_id, err := s.repository.GetClientUUIDByClientID(ctx, req.ClientID)
	if err != nil {
		return nil, errors.New("invalid client ID")
	}
	//verify client id
	auth_request, err := s.repository.GetAuthRequestByClientID(ctx, client_id)
	if err != nil {
		return nil, errors.New("invalid client ID")
	}

	if auth_request.IsExpired() {
		return nil, errors.New("auth request is expired")
	}

	//verify authorization code
	code, err := s.repository.GetAuthCodeByCode(ctx, req.Code)
	if err != nil {
		return nil, errors.New("invalid authorization code")
	}
	if code.IsExpired() {
		return nil, errors.New("authorization code is expired")
	}

	//verify code challenge
	hashedCodeVerifier := hashutil.HashedCodeVerifier(req.CodeVerifier)
	if hashedCodeVerifier != auth_request.CodeChallenge {
		return nil, errors.New(req.CodeVerifier + " does not match the code challenge:" + hashedCodeVerifier + " != " + auth_request.CodeChallenge)
	}

	accesstoken, err := tokenutil.GenerateJWTToken(code.Username, 3600)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}
	refreshtoken, err := randomutil.GenerateRandomString(32)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	//generate access token
	response := &TokenResponse{
		AccessToken:  accesstoken, // Generate a random access token
		TokenType:    "Bearer",
		ExpiresIn:    3600,         // Set token expiration time (1 hour)
		RefreshToken: refreshtoken, // Generate a random refresh token
		Nonce:        req.Nonce,
	}
	return response, nil
}
