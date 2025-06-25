package models

import (
	"encoding/json"
	"sso-go-gin/pkg/utils/hashutil"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;"`
	Username string    `json:"username" gorm:"uniqueIndex;not null"`
	Password string    `json:"password" gorm:"not null"`
	Email    string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
}

func (User) TableName() string {
	return "user_info"
}

func (u *User) CheckPassword(password string) bool {
	hash := hashutil.CheckPasswordHash(password, u.Password)
	return hash
}

type AuthRequestCode struct {
	ID                      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	ClientID                uuid.UUID `json:"client_id" gorm:"type:uuid;not null"`
	ResponseType            string    `json:"response_type" gorm:"type:varchar(30);not null"`
	Scope                   string    `json:"scope" gorm:"type:varchar(100)"`
	State                   string    `json:"state" gorm:"type:varchar(40);not null"`
	Nonce                   *string   `json:"nonce" gorm:"type:varchar(40);uniqueIndex"`
	CodeChallenge           string    `json:"code_challenge" gorm:"type:varchar(100);uniqueIndex;not null"`
	CodeChallengeMethod     string    `json:"code_challenge_method" gorm:"type:varchar(10);not null"`
	AuthRedirectCallbackURI string    `json:"auth_redirect_callback_uri" gorm:"type:varchar(255)"`
	SSORedirectCallbackURI  string    `json:"sso_redirect_callback_uri" gorm:"type:varchar(255)"`
	ExpiredDatetime         time.Time `json:"expired_datetime" gorm:"not null"`
	CreatedDatetime         time.Time `json:"created_datetime" gorm:"not null;default:now()"`
}

func (a *AuthRequestCode) IsExpired() bool {
	//return time.Now().After(a.ExpiredDatetime)
	return false //for testing purposes
}

func (a *AuthRequestCode) BeforeCreate(tx *gorm.DB) (err error) {
	if a.Nonce != nil && *a.Nonce == "" {
		a.Nonce = nil
	}
	return nil
}

func (a *AuthRequestCode) AfterCreate(tx *gorm.DB) (err error) {
	if a.Nonce != nil && *a.Nonce == "" {
		a.Nonce = nil
	}
	return nil
}

type AuthCode struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Code            string    `json:"code" gorm:"type:varchar(100);uniqueIndex;not null"`
	RID             uuid.UUID `json:"rid" gorm:"type:uuid;not null;column:rid"`
	Type            string    `json:"type" gorm:"type:varchar(10);not null"`
	ExpiredDatetime time.Time `json:"expired_datetime" gorm:"not null"`
	CreatedDatetime time.Time `json:"created_datetime" gorm:"not null;default:now()"`
	Username        string    `json:"username" gorm:"type:varchar(50);not null"`
}

func (a *AuthCode) IsExpired() bool {
	// return time.Now().After(a.ExpiredDatetime)
	return false //for testing purposes
}

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

type SSOToken struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Token           string    `json:"token" gorm:"type:varchar(100);uniqueIndex;not null"`
	Source          string    `json:"source" gorm:"type:varchar(50);not null"`
	Destination     string    `json:"destination" gorm:"type:varchar(50);not null"`
	ClientID        uuid.UUID `json:"client_id" gorm:"type:uuid;not null"`
	User            string    `json:"user" gorm:"type:varchar(50);not null"`
	ExpiredDatetime time.Time `json:"expired_datetime" gorm:"not null"`
	CreatedDatetime time.Time `json:"created_datetime" gorm:"not null;default:now()"`
}

func (s *SSOToken) IsExpired() bool {
	// return time.Now().After(s.ExpiredDatetime)
	return false //for testing purposes
}

type SSORequestURI struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	User            string    `json:"user" gorm:"type:varchar(50);not null"`
	ClientID        uuid.UUID `json:"client_id" gorm:"type:uuid;not null"`
	RequestURI      string    `json:"request_uri" gorm:"type:varchar(255);not null;uniqueIndex"`
	SSOToken        string    `json:"sso_token" gorm:"type:varchar(100);not null"`
	AuthRequestID   uuid.UUID `json:"auth_request_id" gorm:"type:uuid;not null"`
	ExpiredDatetime time.Time `json:"expired_datetime" gorm:"not null"`
	CreatedDatetime time.Time `json:"created_datetime" gorm:"not null;default:now()"`
}

func (s *SSORequestURI) IsExpired() bool {
	// return time.Now().After(s.ExpiredDatetime)
	return false //for testing purposes
}
