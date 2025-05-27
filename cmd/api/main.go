package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/authorize", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "received",
		})
	})

	router.Run("localhost:8080")
}
