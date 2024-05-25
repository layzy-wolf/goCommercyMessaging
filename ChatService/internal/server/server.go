package server

import (
	pb "ChatService/api/grpc/chat.v1"
	"ChatService/config"
	"ChatService/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Engine(serv *service.Server, cfg *config.Cfg) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", cfg.Chat.Port))
	if err != nil {
		log.Fatalf("E: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, serv)
	log.Printf("Server listen: %v", lis.Addr())

	return s.Serve(lis)
}
