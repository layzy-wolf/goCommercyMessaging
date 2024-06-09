package service

import (
	pb "AccountService/api/account.v1"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Secret = []byte("my-Secret")
)

type Service struct {
	conn *mongo.Client
	pb.UnimplementedAccountServer
}

func New(client *mongo.Client) *Service {

	return &Service{
		conn: client,
	}
}
