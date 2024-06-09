package endpoint

import (
	"ApiGateway/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type AuthEndpoints struct {
	Login  endpoint.Endpoint
	Verify endpoint.Endpoint
}

type LoginRequest struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

type VerifyRequest struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	Valid bool `json:"valid"`
}

// Создание экземпляра AuthEndpoints
func MakeAuthEndpoints(srv service.AuthService) AuthEndpoints {
	return AuthEndpoints{
		Login:  MakeLoginEndpoint(srv),
		Verify: MakeVerifyEndpoint(srv),
	}
}

// Обертка для сервисного слоя авторизации функции Login
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

func MakeVerifyEndpoint(srv service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(VerifyRequest)
		success := srv.Verify(req.Token)
		return VerifyResponse{
			Valid: success,
		}, nil
	}
}
