// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"user-service/pkg/api"
	"user-service/pkg/api/service"
	"user-service/pkg/config"
	"user-service/pkg/repository"
)

// Injectors from wire.go:

func InitializeServe(cfg *config.Config) (*api.Server, error) {
	userRepo := repository.NewUserRepo()
	userServiceServer := service.NewUserService(userRepo)
	server, err := api.NewGrpcServe(cfg, userServiceServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
