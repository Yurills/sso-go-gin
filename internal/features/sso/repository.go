package sso

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserByUsername(c context.Context, username string) (*User, error) {
	var user User
	if err := r.db.WithContext(c).
		Where("username = ?", username).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// authorization code
func (r *repository) SaveAuthRequest(c context.Context, req *AuthRequestCode) error {
	return r.db.WithContext(c).Create(req).Error
}

func (r *repository) GetAuthRequestByID(c context.Context, id string) (*AuthRequestCode, error) {
	var req AuthRequestCode
	if err := r.db.WithContext(c).First(&req, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *repository) SaveAuthCode(c context.Context, code AuthCode) error {
	return r.db.WithContext(c).Create(code).Error
}
