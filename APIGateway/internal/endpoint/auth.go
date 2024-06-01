package endpoint

import (
	"ApiGateway/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type AuthEndpoints struct {
	Login    endpoint.Endpoint
	Register endpoint.Endpoint
}

type LoginRequest struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

type RegisterRequest struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func MakeAuthEndpoints(srv service.AuthService) AuthEndpoints {
	return AuthEndpoints{
		Login:    MakeLoginEndpoint(srv),
		Register: MakeRegisterEndpoint(srv),
	}
}

func MakeLoginEndpoint(srv service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginRequest)
		token, err := srv.Login(req.Token)
		if err != nil {
			return LoginResponse{
				Token: "",
				Error: err.Error(),
			}, err
		}
		return LoginResponse{
			Token: token,
			Error: "",
		}, nil
	}
}

func MakeRegisterEndpoint(srv service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RegisterRequest)
		success, err := srv.Register(req.Token)
		if err != nil {
			return RegisterResponse{
				Success: success,
				Error:   err.Error(),
			}, err
		}
		return RegisterResponse{
			Success: success,
			Error:   "",
		}, nil
	}
}
