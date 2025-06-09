
package authorize

import (

	"sso-go-gin/internal/sso/authorize/handler"
	"sso-go-gin/internal/sso/authorize/repository"
	"sso-go-gin/internal/sso/authorize/service"

	"github.com/google/wire"
)

var Providers = wire.NewSet(
	repository.NewAuthorizeRepository,
	service.NewAuthorizeService,
	handler.NewAuthorizeHandler,
)