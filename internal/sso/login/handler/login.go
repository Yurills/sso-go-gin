package handler

import (
	"net/http"
	"sso-go-gin/internal/sso/login/dtos"
	"sso-go-gin/internal/sso/login/service"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService *service.LoginService
}

func NewLoginHandler(LoginService *service.LoginService) *LoginHandler {
	return &LoginHandler{LoginService}
}

func (h *LoginHandler) PostLogin(c *gin.Context) {
	var req dtos.LoginRequest

	//verify JSON binding
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//call service to handle login
	res, err := h.LoginService.Login(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	callback_uri := res.RedirectURI
	c.JSON(http.StatusOK, gin.H{
		"code":         res.AuthCode,
		"state":        res.State,
		"nonce":        res.Nonce,
		"redirect_uri": callback_uri})

}

func (h *LoginHandler) ContinueLogin(c *gin.Context) {
	res, err := h.LoginService.ContinueLogin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	callback_uri := res.RedirectURI
	c.JSON(http.StatusOK, gin.H{
		"code":         res.AuthCode,
		"state":        res.State,
		"nonce":        res.Nonce,
		"redirect_uri": callback_uri})
}
