//go:build wireinject
// +build wireinject

package main

import (
	"sso-go-gin/internal/features/login"
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
