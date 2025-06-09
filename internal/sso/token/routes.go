package token

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *TokenHandler) {
	rg.POST("/token", h.PostToken)
}
