//go:build wireinject
// +build wireinject

package sso

import (
	"sso-go-gin/config"
	"sso-go-gin/internal/sso/handler"
	"sso-go-gin/internal/sso/repository"
	"sso-go-gin/internal/sso/service"
	"sso-go-gin/pkg/database"

	"github.com/google/wire"
)

func InitializeSSOHandlers(cfg *config.Config) (*handler.SSOHandlers, error) {
	wire.Build(
		database.NewDB,
		repository.NewLoginRepository,
		service.NewLoginService,
		handler.NewLoginHandler,

		repository.NewAuthorizeRepository,
		service.NewAuthorizeService,
		handler.NewAuthorizeHandler,

		wire.Struct(new(handler.SSOHandlers), "*"),
	)
	return nil, nil
}
