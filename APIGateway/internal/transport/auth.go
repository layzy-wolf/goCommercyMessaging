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
	Login  *http.Server
	Verify *http.Server
}

// Декодирование Login запроса
func decodeHTTPLoginRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Кодирование Login ответа
func encodeHTTPLoginResponse(_ context.Context, w net.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Декодирование Verify запроса
func decodeHTTPVerifyRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.Token = r.Header.Get("Authorization")
	return request, nil
}

// Кодирование Verify ответа
func encodeHTTPVerifyResponse(_ context.Context, w net.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// создание экземпляра http обработчика
func NewAuthHandler(cfg config.Cfg) *AuthServer {
	srv := service.NewAuthService(cfg)
	en := endpoints.MakeAuthEndpoints(srv)
	handler := &AuthServer{
		Login:  http.NewServer(en.Login, decodeHTTPLoginRequest, encodeHTTPLoginResponse),
		Verify: http.NewServer(en.Verify, decodeHTTPVerifyRequest, encodeHTTPVerifyResponse),
	}
	return handler
}
