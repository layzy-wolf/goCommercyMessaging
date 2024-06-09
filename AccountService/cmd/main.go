package main

import (
	"AccountService/config"
	"AccountService/internal/server"
	"AccountService/internal/service"
	"AccountService/internal/storage"
	"log"
)

func main() {
	cfg := config.Load()

	store := storage.New(&cfg)

	serv := service.New(store)

	if err := server.Engine(&cfg, serv); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
