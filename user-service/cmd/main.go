package main

import (
	"log"
	"user-service/pkg/config"
	"user-service/pkg/wire"
)

func main() {
	cfg, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Failed to load config", configErr.Error())
	}

	server, err := wire.InitializeServe(&cfg)
	if err != nil {
		log.Fatal("Failed to init server", err.Error())
	}
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server")
	}

}
