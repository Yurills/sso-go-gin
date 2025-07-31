//go:build wireinject
// +build wireinject

package middleware

import (
	"sso-go-gin/config"
	"sso-go-gin/internal/middleware/protect"
	"sso-go-gin/internal/middleware/protect/handler"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type Middlewares struct {
	AdminOnlyMiddleware *handler.ProtectMiddleware
}

func InitializeMiddlewares(cfg *config.Config, db *gorm.DB) *Middlewares {
	wire.Build(
		protect.ProviderSet,
		wire.Struct(new(Middlewares), "*"),
	)
	return nil
}
