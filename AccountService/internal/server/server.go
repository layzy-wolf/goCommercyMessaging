package server

import (
	pb "AccountService/api/account.v1"
	"AccountService/config"
	"AccountService/internal/service"
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
	pb.RegisterAccountServer(s, srv)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}
