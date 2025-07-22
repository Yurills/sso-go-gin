package logout

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService *LogoutService
}

func NewLogoutHandler(logoutService *LogoutService) *LogoutHandler {
	return &LogoutHandler{LogoutService: logoutService}
}

func (h *LogoutHandler) PostLogout(c *gin.Context) {
	//get username from request form body

	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
		return
	}

	// Call the logout service to handle the logout logic
	if err := h.LogoutService.Logout(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log out"})
		println("Logout error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
