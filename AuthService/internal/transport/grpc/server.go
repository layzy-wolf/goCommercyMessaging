package grpc

import (
	"fmt"
	pb "github.com/layzy-wolf/goCommercyMessaging/AuthService/api/auth"
	"github.com/layzy-wolf/goCommercyMessaging/AuthService/internal/config"
	"github.com/layzy-wolf/goCommercyMessaging/AuthService/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Engine(cfg *config.Cfg, server *service.Server) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPConfig.Port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}
