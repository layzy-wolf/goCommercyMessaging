package grpc

import (
	authv1 "app/api/protobuf/auth.v1"
	"app/internal/config"
	"app/internal/endpoints"
	"context"
	"github.com/go-kit/kit/log/level"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/mhchlib/go-kit/log"
	"google.golang.org/grpc"
	"net"
	"os"
)

type gRPCServer struct {
	authv1.UnimplementedAuthServer
	login    gt.Handler
	register gt.Handler
}

func Engine(cfg *config.Cfg, logger log.Logger, grpcServer authv1.AuthServer) {
	grpcListen, err := net.Listen("tcp", cfg.GRPConfig.Port)

	if err != nil {
		_ = logger.Log("during, Listen", "err", err)
		os.Exit(1)
	}

	baseServer := grpc.NewServer()
	authv1.RegisterAuthServer(baseServer, grpcServer)
	_ = level.Info(logger).Log("msg", "server started!")
	_ = baseServer.Serve(grpcListen)
}

func NewGRPCServer(endpoints endpoints.Endpoints) authv1.AuthServer {
	return &gRPCServer{
		login: gt.NewServer(
			endpoints.Login,
			decodeLoginRequest,
			encodeLoginResponse,
		),
		register: gt.NewServer(
			endpoints.Register,
			decodeRegisterRequest,
			encodeRegisterResponse,
		),
	}
}

func (s *gRPCServer) Login(ctx context.Context, req *authv1.LoginRequest) (res *authv1.LoginResponse, err error) {
	_, resp, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*authv1.LoginResponse), nil
}

func decodeLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*authv1.LoginRequest)
	return endpoints.LoginReq{
		Email:  req.Email,
		Passwd: req.Passwd,
	}, nil
}

func encodeLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.LoginRes)
	return &authv1.LoginResponse{Token: res.Jwt}, nil
}

func (s *gRPCServer) Register(ctx context.Context, req *authv1.RegisterRequest) (res *authv1.RegisterResponse, err error) {
	_, resp, err := s.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*authv1.RegisterResponse), nil
}

func decodeRegisterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*authv1.RegisterRequest)
	return endpoints.RegReq{
		Email:  req.Email,
		Login:  req.Login,
		Passwd: req.Passwd,
	}, nil
}

func encodeRegisterResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.RegRes)
	return &authv1.RegisterResponse{
		UserId: res.UserId,
	}, nil
}
