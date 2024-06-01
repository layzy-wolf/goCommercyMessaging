package main

import (
	"ChatService/config"
	"ChatService/internal/server"
	"ChatService/internal/store"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	cfg := config.Load()

	c := store.NewStore(&cfg)

	defer func(c *mongo.Client, ctx context.Context) {
		err := c.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(c, context.Background())

	s := store.New(c)

	if err := server.Store(s, &cfg); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
