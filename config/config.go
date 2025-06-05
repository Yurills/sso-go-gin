package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Hostname string
	Username string
	Password string
	Port     string
	DBName   string
}

var AppConfig *Config

func Load() *Config {
	_ = godotenv.Load()

	AppConfig = &Config{
		Hostname: getEnv("DB_HOST", "localhost"),
		Username: getEnv("DB_USER", "user"),
		Password: getEnv("DB_PASSWORD", "12345"),
		Port:     getEnv("DB_PORT", "5432"),
		DBName:   getEnv("DB_NAME", "mydatabase"),
	}

	return AppConfig
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
