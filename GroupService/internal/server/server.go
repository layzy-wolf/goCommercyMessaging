package server

import (
	pb "GroupService/api/group.v1"
	"GroupService/config"
	"GroupService/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Engine(cfg *config.Cfg, srv *service.Service) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Group.Port))
	if err != nil {
		log.Fatalf("E: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGroupServer(s, srv)
	log.Printf("Server listening at %v", lis.Addr())
	return s.Serve(lis)
}
