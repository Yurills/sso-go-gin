package token

import (
	"context"
	"sso-go-gin/internal/sso/models"

	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db}
}

func (r *TokenRepository) GetClientUUIDByClientID(c context.Context, client_id string) (string, error) {
	var client models.AuthClient
	if err := r.db.WithContext(c).First(&client, "client_id = ?", client_id).Error; err != nil {
		return "", err
	}
	return client.ID.String(), nil
}

func (r *TokenRepository) GetAuthRequestByID(c context.Context, id string) (*models.AuthRequestCode, error) {
	var req models.AuthRequestCode
	if err := r.db.WithContext(c).First(&req, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *TokenRepository) GetAuthCodeByCode(c context.Context, code string) (*models.AuthCode, error) {
	var authCode models.AuthCode
	if err := r.db.WithContext(c).First(&authCode, "code = ?", code).Error; err != nil {
		return nil, err
	}
	return &authCode, nil
}

func (r *TokenRepository) GetUserByUsername(c context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(c).First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *TokenRepository) GetSSOTokenByClientID(c context.Context, id string) (*models.SSOToken, error) {
	var ssoToken models.SSOToken
	if err := r.db.WithContext(c).First(&ssoToken, "client_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ssoToken, nil
}

func (r *TokenRepository) SaveRefreshToken(c context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(c).Create(token).Error
}

func (r *TokenRepository) GetRefreshToken(c context.Context, token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := r.db.WithContext(c).First(&refreshToken, "refresh_token = ?", token).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *TokenRepository) UpdateRefreshToken(c context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(c).Save(token).Error
}
