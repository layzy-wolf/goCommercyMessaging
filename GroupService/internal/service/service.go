package service

import (
	pb "GroupService/api/group.v1"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	secret = []byte("my-Secret")
)

type Service struct {
	client *mongo.Client
	pb.UnimplementedGroupServer
}

func New(client *mongo.Client) *Service {
	return &Service{client: client}
}
