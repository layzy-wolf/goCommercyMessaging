package endpoints

import (
	"app/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Login    endpoint.Endpoint
	Register endpoint.Endpoint
}

type LoginReq struct {
	Email  string
	Passwd string
}

type LoginRes struct {
	Jwt string
}

type RegReq struct {
	Email  string
	Login  string
	Passwd string
}

type RegRes struct {
	UserId int32
}

func MakeEndpoints(s *service.Service) Endpoints {
	return Endpoints{
		makeLoginEndpoint(s),
		makeRegisterEndpoint(s),
	}
}

func makeLoginEndpoint(s *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginReq)
		res, _ := s.LoginUser(req.Email, req.Passwd)
		return LoginRes{res}, nil
	}
}

func makeRegisterEndpoint(s *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RegReq)
		res, _ := s.RegisterUser(req.Email, req.Login, req.Passwd)
		return RegRes{res}, nil
	}
}
