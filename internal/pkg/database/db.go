package database

import (
	"fmt"
	"sso-go-gin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	config.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		config.AppConfig.Hostname,
		config.AppConfig.Username,
		config.AppConfig.Password,
		config.AppConfig.Port,
		config.AppConfig.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
