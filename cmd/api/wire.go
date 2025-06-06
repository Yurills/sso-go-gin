//go:build wireinject
// +build wireinject

package main

import (
	"sso-go-gin/config"
	"sso-go-gin/internal/sso"
	"sso-go-gin/internal/sso/handler"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*gin.Engine, error) {
	wire.Build(
		sso.InitializeSSOHandlers,
		newRouter,
	)
	return nil, nil
}

func newRouter(h *handler.SSOHandlers) *gin.Engine {
	r := gin.Default()

	ssoGroup := r.Group("/sso")
	handler.RegisterSSORoutes(ssoGroup, h)

	return r
}
