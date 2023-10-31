//go:build wireinject
// +build wireinject

package wire

import (
	"api-gateway/pkg/api"
	"api-gateway/pkg/api/handlers"
	"api-gateway/pkg/client"
	"api-gateway/pkg/config"
	"api-gateway/pkg/service"

	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (*api.Server, error) {
	wire.Build(service.InitClient,
		client.NewUserClient,
		handlers.NewUserHandler,
		api.NewServerHTTP)
	return &api.Server{}, nil
}
