package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *PARHandler) {
	rg.POST("/par/request.token", h.PostRequestToken)
	rg.POST("/par", h.PostPARRequest)
}
