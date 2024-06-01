package service

import (
	pb "ChatService/api/grpc/chatStore.v1"
	"ChatService/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func StoreClient(cfg config.Cfg) (pb.StoreClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", cfg.Store.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("E: %v", err)
	}

	c := pb.NewStoreClient(conn)

	return c, conn
}
