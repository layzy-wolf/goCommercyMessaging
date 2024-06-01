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
	cfg := config.Load()

	c := storage.New(&cfg)

	defer func(c *mongo.Client, ctx context.Context) {
		err := c.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(c, context.Background())

	srv := service.New(c)

	if err := server.Engine(&cfg, srv); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
