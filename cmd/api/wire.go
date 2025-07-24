//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"path/filepath"
	"sso-go-gin/config"
	"sso-go-gin/internal/admin"
	"sso-go-gin/internal/middleware"
	"sso-go-gin/internal/sso"
	authorizeHandler "sso-go-gin/internal/sso/authorize/handler"
	loginHandler "sso-go-gin/internal/sso/login/handler"
	parHandler "sso-go-gin/internal/sso/par/handler"
	tokenHandler "sso-go-gin/internal/sso/token"

	registerHandler "sso-go-gin/internal/admin/register_client/handler"

	"strings"

	"sso-go-gin/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*gin.Engine, error) {
	wire.Build(
		database.NewDB,
		// database.NewRedisClient,
		sso.InitializeSSOHandlers,
		admin.InitializeAdminHandlers,
		middleware.InitializeMiddlewares,
		newRouter,
	)
	return nil, nil
}

func newRouter(
	h *sso.SSOHandlers,
	adminHandler *admin.AdminHandlers,
	middleware *middleware.Middlewares) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New((cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://localhost:8082"}, // Replace with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Important if you're using cookies
	})))

	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "default", "12345", []byte("secret")) // Replace with your Redis configuration
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600, // Set session expiration time (1 hour)
		HttpOnly: true, // Prevent JavaScript access to session cookies
		Secure:   true,
		SameSite: http.SameSiteLaxMode, // Adjust SameSite policy as needed
	})

	r.Use(sessions.Sessions("sso_session", store))

	ssoGroup := r.Group("/api/sso")
	loginHandler.RegisterRoutes(ssoGroup, h.LoginHandler)
	authorizeHandler.RegisterRoutes(ssoGroup, h.AuthorizeHandler)
	tokenHandler.RegisterRoutes(ssoGroup, h.TokenHandler)
	parHandler.RegisterRoutes(ssoGroup, h.PARHandler)

	adminGroup := r.Group("/api/admin", middleware.AdminOnlyMiddleware.AdminOnlyMiddleware())
	registerHandler.RegisterRoutes(adminGroup, adminHandler.RegisterHandler)

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
