//go:build wireinject
// +build wireinject

package main

import (
	"sso-go-gin/config"
	"sso-go-gin/internal/features/auth/login"
	"sso-go-gin/internal/features/auth/register"
	"sso-go-gin/internal/features/sso"
	"sso-go-gin/internal/pkg/database"

	"github.com/google/wire"
)

func InitializeLoginHandler(cfg *config.Config) (*login.Handler, error) {
	wire.Build(
		database.NewDB,
		login.NewRepository,
		login.NewService,
		login.NewHandler)
	return &login.Handler{}, nil
}

func InitializeRegisterHandler(cfg *config.Config) (*register.Handler, error) {
	wire.Build(
		database.NewDB,
		register.NewRepository,
		register.NewService,
		register.NewHandler)
	return &register.Handler{}, nil
}

func initializeSSOHandler(cfg *config.Config) (*sso.Handler, error) {
	wire.Build(
		database.NewDB,
		sso.NewRepository,
		sso.NewService,
		sso.NewHandler)
	return &sso.Handler{}, nil

}
