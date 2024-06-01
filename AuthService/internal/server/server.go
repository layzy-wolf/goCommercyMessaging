package server

import (
	pb "AuthService/api/auth.v1"
	"AuthService/config"
	"AuthService/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Engine(cfg *config.Cfg, srv *service.Service) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPConfig.Port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, srv)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}
