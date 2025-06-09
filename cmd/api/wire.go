//go:build wireinject
// +build wireinject

package main

import (
	"sso-go-gin/config"
	"sso-go-gin/internal/sso"
	authorizeHandler "sso-go-gin/internal/sso/authorize/handler"
	loginHandler "sso-go-gin/internal/sso/login/handler"

	"sso-go-gin/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*gin.Engine, error) {
	wire.Build(
		database.NewDB,
		sso.InitializeSSOHandlers,
		newRouter,
	)
	return nil, nil
}

func newRouter(h *sso.SSOHandlers) *gin.Engine {
	r := gin.Default()

	ssoGroup := r.Group("/sso")
	loginHandler.RegisterRoutes(ssoGroup, h.LoginHandler)
	authorizeHandler.RegisterRoutes(ssoGroup, h.AuthorizeHandler)

	return r
}
