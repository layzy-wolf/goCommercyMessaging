package service

import (
	pb "AuthService/api/auth.v1"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	c      *mongo.Client
	Secret = []byte("my-Secret")
)

type Service struct {
	pb.UnimplementedAuthServer
}

func New(client *mongo.Client) *Service {
	c = client
	return &Service{}
}
