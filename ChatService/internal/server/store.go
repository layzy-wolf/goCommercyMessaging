package server

import (
	pb "ChatService/api/grpc/chatStore.v1"
	"ChatService/config"
	"ChatService/internal/store"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Store(store *store.Server, cfg *config.Cfg) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", cfg.Store.Port))
	if err != nil {
		log.Fatalf("E: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStoreServer(s, store)
	log.Printf("Store Server listen: %v", lis.Addr())

	return s.Serve(lis)
}
