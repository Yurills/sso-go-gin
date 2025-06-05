package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(Service *Service) *Handler {
	return &Handler{Service}
}

func (h *Handler) PostLogin(c *gin.Context) {
	var form LoginRequest

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
	}

	user, err := h.Service.Login(c.Request.Context(), form)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login correct", "user_id": user.ID})

}
