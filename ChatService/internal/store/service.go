package store

import (
	pb "ChatService/api/grpc/chatStore.v1"
	"ChatService/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var c *mongo.Client

type Server struct {
	pb.UnimplementedStoreServer
}

func New(client *mongo.Client) *Server {
	c = client
	s := &Server{}
	return s
}
func NewStore(cfg *config.Cfg) *mongo.Client {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%v", cfg.Store.MongoHost)))
	if err != nil {
		log.Fatalf("E: %v", err)
	}

	log.Print("connected to MongoDB")
	return client
}
