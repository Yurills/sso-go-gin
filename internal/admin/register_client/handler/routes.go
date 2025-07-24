package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *RegisterHandler) {
	r.POST("/register-client", h.RegisterClient)
}
