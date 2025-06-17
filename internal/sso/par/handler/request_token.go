package handler

import (
	"sso-go-gin/internal/sso/par/dtos"
	

	"github.com/gin-gonic/gin"
)



func (h *PARHandler) PostRequestToken(c *gin.Context) {
	// Call the service to handle the request token
	var req dtos.PARRequestTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	//call service to handle request token
	response, err := h.Service.GetRequestToken(c, &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"token": response.Token,
	})
}
