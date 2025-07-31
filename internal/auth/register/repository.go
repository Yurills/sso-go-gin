package register

import (
	"context"
	"sso-go-gin/internal/auth/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(c context.Context, user *models.User) (*models.User, error) {
	if err := r.db.WithContext(c).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetUserInfo(c context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(c).
		Where("username = ?", username).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
