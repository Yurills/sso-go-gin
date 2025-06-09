package repository

import (
	"context"
	"sso-go-gin/internal/sso/models"

	"gorm.io/gorm"
)

type AuthorizeRepository struct {
	db *gorm.DB
}

func NewAuthorizeRepository(db *gorm.DB) *AuthorizeRepository {
	return &AuthorizeRepository{db}
}

func (r *AuthorizeRepository) GetAuthClientByID(c context.Context, client_id string) (*models.AuthClient, error) {
	var req models.AuthClient
	if err := r.db.WithContext(c).First(&req, "client_id = ?", client_id).Error; err != nil {
		return nil, err
	}
	return &req, nil

}

func (r *AuthorizeRepository) SaveRequestCode(c context.Context, code *models.AuthRequestCode) error {
	return r.db.WithContext(c).Create(code).Error
}
