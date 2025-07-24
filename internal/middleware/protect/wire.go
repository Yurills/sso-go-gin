package protect

import (
	"sso-go-gin/internal/middleware/protect/handler"
	"sso-go-gin/internal/middleware/protect/repository"
	"sso-go-gin/internal/middleware/protect/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repository.NewProtectRepository,
	service.NewProtectService,
	handler.NewProtectMiddleware,
)
