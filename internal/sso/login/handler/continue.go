package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
