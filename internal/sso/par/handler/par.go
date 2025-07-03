package handler

import (
	"sso-go-gin/internal/sso/par/dtos"
	"sso-go-gin/internal/sso/par/service"

	"github.com/gin-gonic/gin"
)

type PARHandler struct {
	Service *service.PARService
}

func NewPARHandler(service *service.PARService) *PARHandler {
	return &PARHandler{Service: service}
}

func (h *PARHandler) PostPARRequest(c *gin.Context) {
	// Call the service to handle the PAR request
	var req dtos.PARRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	//if session is active, skip request uri and send token directly
	//TODO: recheck if it should be changed to refresh token
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		print("Error retrieving session ID from cookie:", err)
	} else {
		print("Session ID from cookie:", sessionID)
	}
	if sessionID != "" {
		// Session is active, generate auth code directly
		print("Session is active, generating refresh token directly\n")
		token, err := h.Service.GenerateRefreshToken(c, sessionID, req)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"token": token.RefreshToken,
			"state": token.State,
		})
		return
	}

	response, err := h.Service.CreateRequestURI(c, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"request_uri": response.RequestURI})
}
