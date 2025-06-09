package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *LoginHandler) {
	rg.POST("/login", h.PostLogin)

}
