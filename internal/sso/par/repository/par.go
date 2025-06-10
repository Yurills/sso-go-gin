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

func (r *PARRepository) SaveSSOToken(c context.Context, token *models.SSOToken) error {
	return r.db.WithContext(c).Create(token).Error
}
