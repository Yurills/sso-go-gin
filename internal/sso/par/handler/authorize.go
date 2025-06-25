package handler

import (
	"net/http"
	"sso-go-gin/internal/sso/par/dtos"

	"github.com/gin-gonic/gin"
)

func (h *PARHandler) PostAuthorize(c *gin.Context) {
	var req dtos.PARRequestAuthorize
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	response, err := h.Service.Authorize(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"redirect_uri": response.RedirectURI})
}
