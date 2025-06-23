package token

import "github.com/gin-gonic/gin"

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
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service to handle token generation
	response, err := h.TokenService.GenerateToken(c, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
		"expires_in":    response.ExpiresIn,
	})

}


