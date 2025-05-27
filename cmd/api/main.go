package main

import (
	"fmt"
	"net/http"
	"sso-go-gin/config"
	"sso-go-gin/internal/features/login"
	"sso-go-gin/internal/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		config.AppConfig.Hostname,
		config.AppConfig.Username,
		config.AppConfig.Password,
		config.AppConfig.Port,
		config.AppConfig.DBName)

	db := database.Init(dsn)

	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "received",
		})
	})

	repo := login.NewRepository(db)
	service := login.NewService(repo)
	handler := login.NewHandler(service)

	router.POST("/login", handler.PostLogin)

	router.Run("localhost:8080")
}
