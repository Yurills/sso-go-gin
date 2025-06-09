package handler

import (
	"net/http"
	"sso-go-gin/internal/sso/authorize/dtos"
	"sso-go-gin/internal/sso/authorize/service"

	"github.com/gin-gonic/gin"
)

type AuthorizeHandler struct {
	AuthorizeService *service.AuthorizeService
}

func NewAuthorizeHandler(AuthorizeService *service.AuthorizeService) *AuthorizeHandler {
	return &AuthorizeHandler{AuthorizeService}
}

func (h *AuthorizeHandler) GetAuthorize(c *gin.Context) {
	var req dtos.AuthorizeRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.AuthorizeService.Authorize(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rid": res.RID,
	})
}
