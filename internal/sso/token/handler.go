package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	TokenService *TokenService
}

func NewTokenHandler(tokenService *TokenService) *TokenHandler {
	return &TokenHandler{
		TokenService: tokenService,
	}
}

func (h *TokenHandler) PostToken(c *gin.Context) {
	var req TokenRequest

	// Verify JSON binding
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service to handle token generation
	response, err := h.TokenService.GenerateToken(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":     response.AccessToken,
		"refresh_token":    response.RefreshToken,
		"expires_in":       response.ExpiresIn,
		"destination_link": response.DestinationLink,
	})

}
