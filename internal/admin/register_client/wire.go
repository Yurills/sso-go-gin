package register_client

import (
	"sso-go-gin/internal/admin/register_client/handler"
	"sso-go-gin/internal/admin/register_client/repository"
	"sso-go-gin/internal/admin/register_client/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repository.NewRegisterRepository,
	service.NewRegisterService,
	handler.NewRegisterHandler,
)
