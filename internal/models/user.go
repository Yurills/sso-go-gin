package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:gen_random_uuid()"`
	Username string    `json:"username" gorm:"uniqueIndex;not null"`
	Password string    `json:"password" gorm:"not null"`
}

func (User) TableName() string {
	return "user_info"
}
