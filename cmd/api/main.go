package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "received",
		})
	})

	loginHandler, err := InitializeLoginHandler()
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	registerHandler, err := InitializeRegisterHandler()
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	router.POST("/login", loginHandler.PostLogin)
	router.POST("/register", registerHandler.PostRegister)

	router.Run("localhost:8080")
}
