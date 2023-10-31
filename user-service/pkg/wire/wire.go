//go:build wireinject

package wire

import (
	"user-service/pkg/api"
	"user-service/pkg/api/service"
	"user-service/pkg/config"
	"user-service/pkg/repository"

	"github.com/google/wire"
)

func InitializeServe(cfg *config.Config) (*api.Server, error) {
	wire.Build(
		repository.NewUserRepo,
		service.NewUserService,
		api.NewGrpcServe)
	return &api.Server{}, nil
}
