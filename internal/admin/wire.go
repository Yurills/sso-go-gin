//go:build wireinject
// +build wireinject

package admin

import (
	"sso-go-gin/config"
	"sso-go-gin/internal/admin/register_client"
	"sso-go-gin/internal/admin/register_client/handler"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type AdminHandlers struct {
	RegisterHandler *handler.RegisterHandler
}

func InitializeAdminHandlers(cfg *config.Config, db *gorm.DB) (*AdminHandlers, error) {
	wire.Build(
		register_client.ProviderSet,
		wire.Struct(new(AdminHandlers), "*"),
	)
	return nil, nil
}
