package repository

import (
	"context"
	"sso-go-gin/internal/sso/models"

	"gorm.io/gorm"
)

type ProtectRepository struct {
	db *gorm.DB
}

func NewProtectRepository(db *gorm.DB) *ProtectRepository {
	return &ProtectRepository{db: db}
}

func (r *ProtectRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ProtectRepository) GetUserIDBySessionID(ctx context.Context, sessionID string) (string, error) {
	var session models.Session
	if err := r.db.WithContext(ctx).First(&session, "id = ?", sessionID).Error; err != nil {
		return "", err
	}
	return session.UserID.String(), nil
}
