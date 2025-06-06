package handler

import (
	
	"github.com/gin-gonic/gin"
)

type SSOHandlers struct {
	Login *LoginHandler
	//Authorize *AuthorizeHandler
}

func RegisterSSORoutes(rg *gin.RouterGroup, h* SSOHandlers) {
	rg.POST("/login", h.Login.PostLogin)
	//rg.GET("/authorize", h.Authorize.GetAuthorize)
}
