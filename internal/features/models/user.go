package models

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Password  string `json:"password" gorm:"not null"`
	
}

func (User) TableName() string {
	return "user_info"
}