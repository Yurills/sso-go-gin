package logout

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewLogoutRepository,
	NewLogoutService,
	NewLogoutHandler,
)
