package repository

import (
	"context"
	"sso-go-gin/internal/client/models"

	"gorm.io/gorm"
)

type RegisterRepository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) *RegisterRepository {
	return &RegisterRepository{
		db: db,
	}
}

func (r *RegisterRepository) CreateClient(c context.Context, client *models.AuthClient) error {
	return r.db.WithContext(c).Create(client).Error
}

func (r *RegisterRepository) GetClientByID(c context.Context, id string) (*models.AuthClient, error) {
	var client models.AuthClient
	err := r.db.WithContext(c).Where("id = ?", id).First(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}
