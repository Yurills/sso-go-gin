package database

import (
	"fmt"
	"sso-go-gin/config"
	"sso-go-gin/internal/features/sso"
	"sso-go-gin/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {


	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		config.Hostname,
		config.Username,
		config.Password,
		config.Port,
		config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// db.AutoMigrate(
	// 	&models.User{},
	// 	&sso.AuthRequestCode{},
	// 	&sso.AuthCode{},
	// )

	return db, nil

}
