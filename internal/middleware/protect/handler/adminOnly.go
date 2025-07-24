package handler

import (
	"net/http"
	"sso-go-gin/internal/middleware/protect/service"

	"github.com/gin-gonic/gin"
)

type ProtectMiddleware struct {
	service *service.ProtectService
}

func NewProtectMiddleware(service *service.ProtectService) *ProtectMiddleware {
	return &ProtectMiddleware{service: service}
}

func (m *ProtectMiddleware) AdminOnlyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionID, err := ctx.Cookie("session_id")
		if err != nil || sessionID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Session ID is required"})
			ctx.Abort()
			return
		}

		user, err := m.service.ProtectAdmin(ctx, sessionID)

		if err != nil || user == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
