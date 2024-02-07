package cmd

import (
	"github.com/layzy-wolf/goCommercyMessaging/AuthService/internal/config"
	"github.com/layzy-wolf/goCommercyMessaging/AuthService/internal/service"
	"github.com/layzy-wolf/goCommercyMessaging/AuthService/internal/storage"
	"github.com/layzy-wolf/goCommercyMessaging/AuthService/internal/transport/grpc"
	"log"
)

func main() {
	cfg := config.MustLoad()

	store := storage.New(cfg.Storage, cfg.StoragePort, cfg.StorageLogin, cfg.StoragePass)

	s := service.New(store)

	if err := grpc.Engine(cfg, s); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
