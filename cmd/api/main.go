package main

import (
	"log"
	"net/http"

	"sso-go-gin/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "received",
		})
	})

	handler, err := InitializeLoginHandler()
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	router.POST("/login", handler.PostLogin)

	router.Run("localhost:8080")
}
