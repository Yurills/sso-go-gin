package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type AuthClient struct {
	ID                      uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name                    string          `gorm:"type:varchar(50);not null" json:"name"`
	Description             string          `gorm:"type:text;not null" json:"description"`
	ClientID                string          `gorm:"type:varchar(100);uniqueIndex;not null" json:"client_id"`
	ClientSecret            string          `gorm:"type:varchar(100);uniqueIndex;not null" json:"client_secret"`
	AuthRedirectCallbackURI string          `gorm:"type:varchar(255);uniqueIndex;not null" json:"auth_redirect_callback_uri"`
	SSORedirectCallbackURI  string          `gorm:"type:varchar(255);uniqueIndex;not null" json:"sso_redirect_callback_uri"`
	Scope                   string          `gorm:"type:varchar(100)" json:"scope"`
	Active                  bool            `gorm:"not null;default:true" json:"active"`
	ConfigProfile           json.RawMessage `gorm:"type:jsonb;not null;default:'{}'" json:"config_profile"`
	PrivateKey              string          `gorm:"type:text;not null" json:"private_key"`
	PublicKey               string          `gorm:"type:text;not null" json:"public_key"`
	CreatedDatetime         time.Time       `gorm:"not null;default:now()" json:"created_datetime"`
	UpdatedDatetime         time.Time       `gorm:"not null" json:"updated_datetime"`
}
