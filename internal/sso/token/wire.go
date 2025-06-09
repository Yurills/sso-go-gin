package token

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewTokenRepository,
	NewTokenService,
	NewTokenHandler,
)
