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

func (r *TokenRepository) GetAuthRequestByClientID(c context.Context, client_id string) (*models.AuthRequestCode, error) {
	var req models.AuthRequestCode
	if err := r.db.First(&req, "client_id = ?", client_id).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *TokenRepository) GetAuthCodeByCode(c context.Context, code string) (*models.AuthCode, error) {
	var authCode models.AuthCode
	if err := r.db.First(&authCode, "code = ?", code).Error; err != nil {
		return nil, err
	}
	return &authCode, nil
}
