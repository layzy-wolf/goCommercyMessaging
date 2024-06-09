package store

import (
	"GroupService/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func New(cfg *config.Cfg) *mongo.Client {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%v", cfg.MongoDB)))
	if err != nil {
		log.Fatalf("E: %v", err)
	}

	log.Print("Connected to mongoDB")
	return client
}
