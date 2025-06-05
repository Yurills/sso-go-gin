package main

import (
	"log"
	"net/http"
	"sso-go-gin/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "received",
		})
	})

	loginHandler, err := InitializeLoginHandler(cfg)
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	registerHandler, err := InitializeRegisterHandler(cfg)
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	SSOHandler, err := initializeSSOHandler(cfg)
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	router.POST("/login", loginHandler.PostLogin)
	router.POST("/register", registerHandler.PostRegister)
	router.POST("/sso/login", SSOHandler.PostLogin)

	router.Run("localhost:8080")
}
