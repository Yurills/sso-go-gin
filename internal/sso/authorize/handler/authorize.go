package handler

import (
	"fmt"
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

	// If the user is already logged in, redirect to the callback URL
	sessionID, err := c.Cookie("session_id")
	print("Session ID from cookie:", sessionID, "Error:", err)
	if err != nil {
		session := h.AuthorizeService.ValidSession(c, sessionID)
		if !session {
			// no session, redirect to the login
			res, err := h.AuthorizeService.Authorize(c, req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			LoginURL := fmt.Sprintf("/login?rid=%s&client_id=%s", res.RID, req.ClientID)
			c.Redirect(http.StatusFound, LoginURL)
			return
		}
	}

	res, err := h.AuthorizeService.GenerateAuthCode(c, sessionID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	callbackurl := res.RedirectURI + "?code=" + res.AuthCode + "&state=" + res.State
	if res.Nonce != nil && *res.Nonce != "" {
		callbackurl += "&nonce=" + *res.Nonce
	}
	c.Redirect(http.StatusFound, callbackurl)

}
