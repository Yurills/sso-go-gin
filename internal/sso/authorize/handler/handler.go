package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *AuthorizeHandler) {
		rg.GET("/authorize", h.GetAuthorize)
}