package repository

import (
	"context"
	"sso-go-gin/internal/sso/models"

	"gorm.io/gorm"
)

type PARRepository struct {
	db *gorm.DB
}

func NewPARRepository(db *gorm.DB) *PARRepository {
	return &PARRepository{db: db}
}

func (r *PARRepository) GetAuthRequestByClientID(c context.Context, clientID string) (*models.AuthRequestCode, error) {
	var authRequest models.AuthRequestCode
	if err := r.db.WithContext(c).Where("client_id = ?", clientID).First(&authRequest).Error; err != nil {
		return nil, err
	}
	return &authRequest, nil
}

func (r *PARRepository) GetAuthClientByID(c context.Context, clientID string) (*models.AuthClient, error) {
	var authClient models.AuthClient
	if err := r.db.WithContext(c).Where("client_id = ?", clientID).First(&authClient).Error; err != nil {
		return nil, err
	}
	return &authClient, nil
}

func (r *PARRepository) GetSSOTokenByClientID(c context.Context, clientID string) (*models.SSOToken, error) {
	var ssoToken models.SSOToken
	if err := r.db.WithContext(c).Where("client_id = ?", clientID).First(&ssoToken).Error; err != nil {
		return nil, err
	}
	return &ssoToken, nil
}

func (r *PARRepository) SaveSSOToken(c context.Context, token *models.SSOToken) error {
	return r.db.WithContext(c).Create(token).Error
}

func (r *PARRepository) SaveAuthRequest(c context.Context, req *models.AuthRequestCode) error {
	return r.db.WithContext(c).Create(req).Error
}

func (r *PARRepository) SaveSSORequestURI(c context.Context, uri *models.SSORequestURI) error {
	return r.db.WithContext(c).Create(uri).Error
}
