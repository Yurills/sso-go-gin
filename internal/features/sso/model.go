package sso

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
}

type AuthRequestCode struct {
	ID          string `json:"id" gorm:"primaryKey"`
	ClientID    string `json:"client_id" gorm:"not null"`
	RedirectURI string `json:"redirect_uri" gorm:"not null"`
	State       string `json:"state" gorm:"not null"`
	Scope       string `json:"scope"`
	ExpiresAt   time.Time `json:"expires_at" gorm:"not null"`
}

func (a *AuthRequestCode) IsExpired() bool {
	return time.Now().After(a.ExpiresAt)
}


type AuthCode struct {
	Code      string `json:"code" gorm:"primaryKey"`
	ClientID  string `json:"client_id" gorm:"not null"`
	Type      string `json:"type" gorm:"not null"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
}
