package database

import (
	"fmt"
	"sso-go-gin/config"

	"github.com/redis/go-redis/v9"
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

func NewRedisClient(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
	})
}
