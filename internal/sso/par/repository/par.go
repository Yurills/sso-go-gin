package repository

import (
	"context"
	"sso-go-gin/internal/sso/models"

	"github.com/google/uuid"
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

func (r *PARRepository) SaveAuthRequest(c context.Context, req *models.AuthRequestCode) error {
	return r.db.WithContext(c).Create(req).Error
}

func (r *PARRepository) SaveSSORequestURI(c context.Context, uri *models.SSORequestURI) error {
	return r.db.WithContext(c).Create(uri).Error
}

func (r *PARRepository) GetSSOTokenByToken(c context.Context, tokenID string) (*models.SSOToken, error) {
	var ssoToken models.SSOToken
	if err := r.db.WithContext(c).Where("token = ?", tokenID).First(&ssoToken).Error; err != nil {
		return nil, err
	}
	return &ssoToken, nil
}

func (r *PARRepository) GetAuthClientByName(c context.Context, clientName string) (*models.AuthClient, error) {
	var authClient models.AuthClient
	if err := r.db.WithContext(c).Where("name = ?", clientName).First(&authClient).Error; err != nil {
		return nil, err
	}
	return &authClient, nil
}

func (r *PARRepository) GetUserInfoBySessionID(c context.Context, sessionID string) (*models.User, error) {
	var user models.User
	//join session and user_info table
	session_id, err := uuid.Parse(sessionID)
	if err != nil {
		return nil, err
	}

	if err := r.db.WithContext(c).Table("user_info").Select("user_info.*").
		Joins("JOIN sessions ON user_info.id = sessions.user_id").
		Where("sessions.id = ?", session_id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PARRepository) GetSessionByID(c context.Context, sessionID string) (*models.Session, error) {
	var session models.Session
	if err := r.db.WithContext(c).Where("id = ?", sessionID).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *PARRepository) SaveRefreshToken(c context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(c).Create(token).Error
}
