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

	response, err := h.Service.CreateRequestURI(c, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}
