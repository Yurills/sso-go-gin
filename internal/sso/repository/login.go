package repository
import (
	"context"

	"gorm.io/gorm"

	"sso-go-gin/internal/sso/models"

	
)

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
	return &LoginRepository{db}
}

// authorization code
func (r *LoginRepository) SaveAuthRequest(c context.Context, req *models.AuthRequestCode) error {
	return r.db.WithContext(c).Create(req).Error
}

func (r *LoginRepository) GetAuthRequestByID(c context.Context, id string) (*models.AuthRequestCode, error) {
	var req models.AuthRequestCode
	if err := r.db.WithContext(c).First(&req, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *LoginRepository) SaveAuthCode(c context.Context, code *models.AuthCode) error {
	return r.db.WithContext(c).Create(code).Error
}
