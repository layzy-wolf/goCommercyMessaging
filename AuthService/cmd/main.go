package main

import (
	"app/internal/config"
	endpoints "app/internal/endpoints"
	"app/internal/log"
	"app/internal/service"
	"app/internal/storage"
	"app/internal/transport/grpc"
)

func main() {
	cfg := config.MustLoad()

	logger := log.SetupLogger(cfg.Env)

	store := storage.New(logger, cfg.Storage, cfg.StoragePort, cfg.StorageLogin, cfg.StoragePass)

	s := service.New(store, logger)

	endpoint := endpoints.MakeEndpoints(s)

	grpcServer := grpc.NewGRPCServer(endpoint)

	grpc.Engine(cfg, logger, grpcServer)
}
