package main

import (
	"GroupService/config"
	"GroupService/internal/server"
	"GroupService/internal/service"
	"GroupService/internal/store"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	cfg := config.Load()

	c := store.New(&cfg)

	defer func(c *mongo.Client, ctx context.Context) {
		err := c.Disconnect(ctx)
		if err != nil {
			log.Printf("E: %v", err)
		}
	}(c, context.Background())

	srv := service.New(c)

	if err := server.Engine(&cfg, srv); err != nil {
		log.Fatalf("E: %v", err)
	}
}
