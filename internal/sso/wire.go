//go:build wireinject
// +build wireinject

package sso

import (
	"sso-go-gin/config"
	authorizeHandler "sso-go-gin/internal/sso/authorize/handler"
	loginHandler "sso-go-gin/internal/sso/login/handler"
	"sso-go-gin/internal/sso/token"

	"sso-go-gin/internal/sso/authorize"
	"sso-go-gin/internal/sso/login"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type SSOHandlers struct {
	LoginHandler     *loginHandler.LoginHandler
	AuthorizeHandler *authorizeHandler.AuthorizeHandler
	TokenHandler     *token.TokenHandler
}

func InitializeSSOHandlers(cfg *config.Config, db *gorm.DB) (*SSOHandlers, error) {
	wire.Build(
		login.ProviderSet,
		authorize.Providers,
		token.ProviderSet,
		wire.Struct(new(SSOHandlers), "*"),
	)
	return nil, nil

}
