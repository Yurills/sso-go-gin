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
