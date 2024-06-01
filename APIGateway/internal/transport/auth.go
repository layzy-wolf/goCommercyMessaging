package transport

import (
	"ApiGateway/config"
	endpoints "ApiGateway/internal/endpoint"
	"ApiGateway/internal/service"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/transport/http"
	net "net/http"
)

type AuthServer struct {
	Login    *http.Server
	Register *http.Server
}

func decodeHTTPLoginRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeHTTPLoginResponse(_ context.Context, w net.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTPRegisterRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeHTTPRegisterResponse(_ context.Context, w net.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewAuthHandler(cfg config.Cfg) *AuthServer {
	srv := service.NewAuthService(cfg)
	en := endpoints.MakeAuthEndpoints(srv)
	handler := &AuthServer{
		Login:    http.NewServer(en.Login, decodeHTTPLoginRequest, encodeHTTPLoginResponse),
		Register: http.NewServer(en.Register, decodeHTTPRegisterRequest, encodeHTTPRegisterResponse),
	}
	return handler
}
