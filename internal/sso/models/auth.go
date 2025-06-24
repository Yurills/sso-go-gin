package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CreatedDatetime time.Time `json:"created_datetime" gorm:"not null;default:now()"`
	ExpiredDatetime time.Time `json:"expired_datetime" gorm:"not null"`
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiredDatetime)
}
