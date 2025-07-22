package logout

import (
	"sso-go-gin/internal/sso/models"

	"gorm.io/gorm"
)

type LogoutRepository struct {
	db *gorm.DB
}

func NewLogoutRepository(db *gorm.DB) *LogoutRepository {
	return &LogoutRepository{
		db: db,
	}
}

func (r *LogoutRepository) DeleteSessionByID(sessionID string) error {
	return r.db.Where("id = ?", sessionID).Delete(&models.Session{}).Error
}

func (r *LogoutRepository) GetSessionByUserID(userID string) (*models.Session, error) {
	var session models.Session
	if err := r.db.Where("user_id = ?", userID).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *LogoutRepository) GetUserIDByEmail(email string) (string, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}
	return user.ID.String(), nil
}
