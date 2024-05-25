package main

import (
	"ChatService/config"
	"ChatService/internal/server"
	"ChatService/internal/service"
	"log"
)

func main() {
	cfg := config.Load()

	client, conn := service.StoreClient(cfg)

	defer conn.Close()

	serv := service.New(client)

	if err := server.Engine(serv, &cfg); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
