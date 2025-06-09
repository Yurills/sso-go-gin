package login

import (
	"sso-go-gin/internal/sso/login/handler"
	"sso-go-gin/internal/sso/login/repository"
	"sso-go-gin/internal/sso/login/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repository.NewLoginRepository,
	service.NewLoginService,
	handler.NewLoginHandler,
)
