package main

import (
	"AuthService/config"
	"AuthService/internal/server"
	"AuthService/internal/service"
	"AuthService/internal/storage"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {

	// получение конфига
	cfg := config.Load()

	// инициализация подключения с бд
	c := storage.New(&cfg)

	// определение функции для отключения от бд, при отключении сервиса
	defer func(c *mongo.Client, ctx context.Context) {
		err := c.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(c, context.Background())

	// инициализация сервиса
	srv := service.New(c)

	// запуск сервиса
	if err := server.Engine(&cfg, srv); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
