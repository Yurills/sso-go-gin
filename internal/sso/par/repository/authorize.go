package repository

import (
	"context"
	"sso-go-gin/internal/sso/models"
)

func (r *PARRepository) GetSSOTokenByRequestURI(c context.Context, requestURI string) (*models.SSORequestURI, error) {
	var ssoRequestURI models.SSORequestURI
	if err := r.db.WithContext(c).Where("request_uri = ?", requestURI).First(&ssoRequestURI).Error; err != nil {
		return nil, err
	}
	return &ssoRequestURI, nil
}

func (r *PARRepository) GetAuthRequestByID(c context.Context, id string) (*models.AuthRequestCode, error) {
	var authRequest models.AuthRequestCode
	if err := r.db.WithContext(c).Where("id = ?", id).First(&authRequest).Error; err != nil {
		return nil, err
	}
	return &authRequest, nil
}

func (r *PARRepository) SaveAuthCode(c context.Context, authCode *models.AuthCode) error {
	return r.db.WithContext(c).Create(authCode).Error
}

