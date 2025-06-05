package register

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) PostRegister(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
	}
	// Save the user to the database
	createdUser, err := h.Service.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "user registered successfully", "user_id": createdUser.ID})
}
