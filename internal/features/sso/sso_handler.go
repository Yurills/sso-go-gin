package sso

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(Service *Service) *Handler {
	return &Handler{Service}
}

func (h *Handler) PostLogin(c *gin.Context) {
	var req LoginRequest

	//verify JSON binding
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	//call service to handle login
	res, err := h.Service.Login(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	callback_uri := res.RedirectURI + "?code=" + res.AuthCode + "&state=" + res.State
	c.JSON(http.StatusOK, gin.H{
		"redirect_url": callback_uri})

}
