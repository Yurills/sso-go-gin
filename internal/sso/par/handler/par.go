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
	sessionID, err := c.Cookie("session_id")
	if err == nil && sessionID != "" {
		// Session is active, generate auth code directly
		code, err := h.Service.GenerateAuthCode(c, sessionID, req)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"code": code,
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
