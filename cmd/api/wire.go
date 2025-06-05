//go:build wireinject
// +build wireinject

package main

import (
	"sso-go-gin/internal/features/auth/login"
	"sso-go-gin/internal/features/auth/register"
	"sso-go-gin/internal/pkg/database"

	"github.com/google/wire"
)

func InitializeLoginHandler() (*login.Handler, error) {
	wire.Build(
		database.NewDB,
		login.NewRepository,
		login.NewService,
		login.NewHandler)
	return &login.Handler{}, nil
}

func InitializeRegisterHandler() (*register.Handler, error) {
	wire.Build(
		database.NewDB,
		register.NewRepository,
		register.NewService,
		register.NewHandler)
	return &register.Handler{}, nil
}
