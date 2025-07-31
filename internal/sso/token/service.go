package token

import (
	"errors"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/pkg/utils/hashutil"
	"sso-go-gin/pkg/utils/randomutil"
	"sso-go-gin/pkg/utils/tokenutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TokenService struct {
	repository *TokenRepository
}

func NewTokenService(repository *TokenRepository) *TokenService {
	return &TokenService{repository}
}

// grantType: "authorization_code"
func (s *TokenService) GenerateToken(ctx *gin.Context, req TokenRequest) (*TokenResponse, error) {
	// verify nonce

	// verify grant type

	if req.GrantType != "authorization_code" && req.GrantType != "refresh_token" {
		return nil, errors.New("invalid grant type")
	}

	//refresh access token if refresh token is provided
	if req.GrantType == "refresh_token" {
		if req.Code == "" {
			return nil, errors.New("refresh token is required")
		}
		return s.RefreshToken(ctx, req)

	}
	//verify required parameters
	if req.ClientID == "" || req.Code == "" || req.CodeVerifier == "" {
		return nil, errors.New("client_id, code, and code_verifier are required")
	}

	//verify authorization code
	code, err := s.repository.GetAuthCodeByCode(ctx, req.Code)
	if err != nil {
		return nil, errors.New("invalid authorization code")
	}
	if code.IsExpired() {
		return nil, errors.New("authorization code is expired")
	}

	//verify client id
	authRequest, err := s.repository.GetAuthRequestByID(ctx, code.RID.String())
	if err != nil {
		return nil, errors.New("invalid client ID")
	}
	if authRequest.IsExpired() {
		return nil, errors.New("auth request is expired")
	}

	//verify code challenge
	if !validateCodeChallenge(req.CodeVerifier, authRequest.CodeChallenge) {
		return nil, errors.New("code challenge verification failed")
	}

	//verify user
	user, err := s.repository.GetUserByUsername(ctx, code.Username)
	if err != nil {
		return nil, errors.New("invalid username")
	}

	jwtParams := tokenutil.JWTTokenParams{
		ID:       user.ID.String(),
		Username: code.Username,
		Email:    user.Email,
		Nonce:    authRequest.Nonce,
		TTL:      3600, // Set token expiration time (1 hour)
	}

	accesstoken, err := tokenutil.GenerateJWTToken(jwtParams)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}
	refreshtoken, err := randomutil.GenerateRandomString(32)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// Get destination link from SSO token, client id can have multiple SSO tokens so not really good solution
	// If the SSO token is not found, destination_link will be an empty string
	var destinationLink string
	ssoToken, _ := s.repository.GetSSOTokenByClientID(ctx, authRequest.ClientID.String())
	if ssoToken != nil {
		destinationLink = ssoToken.Destination
	}

	//generate access token
	response := &TokenResponse{
		AccessToken:     accesstoken, // Generate a random access token
		TokenType:       "Bearer",
		ExpiresIn:       3600,         // Set token expiration time (1 hour)
		RefreshToken:    refreshtoken, // Generate a random refresh token
		Nonce:           req.Nonce,
		DestinationLink: destinationLink,
	}

	// Save the SSO token with the access token and refresh token
	refreshToken := models.RefreshToken{
		ID: uuid.New(),
		RefreshToken:    refreshtoken,
		ClientID:        authRequest.ClientID,
		User:            user.Username,
		Email:           user.Email,
		ExpiredDatetime: time.Now().Add(1 * time.Hour), // Set expiration time for the access token
		CreatedDatetime: time.Now(),
	}

	if err := s.repository.SaveRefreshToken(ctx, &refreshToken); err != nil {
		return nil, errors.New("failed to save refresh token")
	}

	return response, nil
}


// grantType: "refresh_token"
func (s *TokenService) RefreshToken(ctx *gin.Context, req TokenRequest) (*TokenResponse, error) {
	// verify grant type
	if req.GrantType != "refresh_token" {
		return nil, errors.New("invalid grant type")
	}

	// verify refresh token
	refreshToken, err := s.repository.GetRefreshToken(ctx, req.Code)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}
	if refreshToken.IsExpired() {
		return nil, errors.New("refresh token is expired")
	}

	// Generate new access token
	jwtParams := tokenutil.JWTTokenParams{
		Username: refreshToken.User,
		Email:    refreshToken.Email,
		Nonce:    &req.Nonce,
		TTL:      3600, // Set token expiration time (1 hour)
	}

	accesstoken, err := tokenutil.GenerateJWTToken(jwtParams)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}
	refreshtoken, err := randomutil.GenerateRandomString(32)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	//generate access token response
	response := &TokenResponse{
		AccessToken:     accesstoken, // Generate a random access token
		TokenType:       "Bearer",
		ExpiresIn:       3600,         // Set token expiration time (1 hour
		RefreshToken:    refreshtoken, // Generate a random refresh token
		Nonce:           req.Nonce,
		DestinationLink: "", // Use the existing destination link from the accessToken
	}

	//set current refresh token to expired
	refreshToken.ExpiredDatetime = time.Now()
	if err := s.repository.UpdateRefreshToken(ctx, refreshToken); err != nil {
		return nil, errors.New("failed to update refresh token")
	}

	return response, nil
}

func validateCodeChallenge(codeVerifier string, codeChallenge string) bool {
	hashedCodeVerifier := hashutil.HashedCodeVerifier(codeVerifier)
	return hashedCodeVerifier == codeChallenge
}
