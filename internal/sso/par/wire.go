package par

import (
	"sso-go-gin/internal/sso/par/handler"
	"sso-go-gin/internal/sso/par/repository"
	"sso-go-gin/internal/sso/par/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repository.NewPARRepository,
	service.NewPARService,
	handler.NewPARHandler,
)
