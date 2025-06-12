//go:build wireinject
// +build wireinject

package main

import (
	"path/filepath"
	"sso-go-gin/config"
	"sso-go-gin/internal/sso"
	authorizeHandler "sso-go-gin/internal/sso/authorize/handler"
	loginHandler "sso-go-gin/internal/sso/login/handler"
	parHandler "sso-go-gin/internal/sso/par/handler"
	tokenHandler "sso-go-gin/internal/sso/token"
	"strings"

	"sso-go-gin/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*gin.Engine, error) {
	wire.Build(
		database.NewDB,
		database.NewRedisClient,
		sso.InitializeSSOHandlers,
		newRouter,
	)
	return nil, nil
}

func newRouter(h *sso.SSOHandlers) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New((cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://localhost:8082"}, // Replace with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Important if you're using cookies
	})))

	ssoGroup := r.Group("/api/sso")
	loginHandler.RegisterRoutes(ssoGroup, h.LoginHandler)
	authorizeHandler.RegisterRoutes(ssoGroup, h.AuthorizeHandler)
	tokenHandler.RegisterRoutes(ssoGroup, h.TokenHandler)
	parHandler.RegisterRoutes(ssoGroup, h.PARHandler)

	staticDir := "./frontend/dist"
	r.Static("/assets", filepath.Join(staticDir, "assets"))

	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(404, gin.H{"error": "Not Found"})
			return
		}

		if strings.Contains(c.Request.URL.Path, ".") {
			c.Status(404)
			return
		}

		c.File(filepath.Join(staticDir, "index.html"))
	})

	return r
}
