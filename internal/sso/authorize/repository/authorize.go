package repository

import (
	"context"
	"sso-go-gin/internal/sso/models"

	"github.com/google/uuid"
	
	"gorm.io/gorm"
)

type AuthorizeRepository struct {
	db    *gorm.DB
}

func NewAuthorizeRepository(db *gorm.DB) *AuthorizeRepository {
	return &AuthorizeRepository{db: db}
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

// func (r *AuthorizeRepository) SaveCSRFToken(c context.Context, csrfToken string, authRequestCodeID string, ttl time.Duration) error {
// 	return r.redis.Set(c, csrfToken, authRequestCodeID, ttl).Err()
// }

// func (r *AuthorizeRepository) GetCSRFToken(c context.Context, csrfToken string) (string, error) {
// 	val, err := r.redis.Get(c, csrfToken).Result()
// 	if err != nil {
// 		if err == redis.Nil {
// 			return "", nil // Token not found
// 		}
// 		return "", err // Other error
// 	}
// 	return val, nil
// }

func (r *AuthorizeRepository) GetSessionByID(c context.Context, sessionID string) (*models.Session, error) {
	var session models.Session
	session_id, err := uuid.Parse(sessionID)
	if err != nil {
		return nil, err
	}
	if err := r.db.WithContext(c).First(&session, "id = ?", session_id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *AuthorizeRepository) GetUserInfoBySessionID(c context.Context, sessionID string) (*models.User, error) {
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
func (r *AuthorizeRepository) SaveAuthCode(c context.Context, authCode *models.AuthCode) error {
	return r.db.WithContext(c).Create(authCode).Error
}
