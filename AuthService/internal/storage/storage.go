package storage

import (
	"AuthService/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func New(cfg *config.Cfg) *mongo.Client {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%v", cfg.MongoHost)))
	if err != nil {
		panic("failed to connect with db")
	}

	log.Print("connected to MongoDB")
	return client
}
