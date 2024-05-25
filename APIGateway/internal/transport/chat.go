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

type ChatServer struct {
	Show  *http.Server
	Chats *http.Server
}

func decodeHTTPShowRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.ShowRequest
	ok, usr, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.From = usr
	return request, nil
}

func decodeHTTPChatsRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.ChatsRequest
	ok, usr, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.User = usr
	return request, nil
}

func encodeHTTPChatResponse(_ context.Context, w net.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewChatHandler(cfg config.Cfg) *ChatServer {
	srv := service.NewChatService(cfg)
	en := endpoints.MakeChatEndpoints(srv)
	handler := &ChatServer{
		Show:  http.NewServer(en.Show, decodeHTTPShowRequest, encodeHTTPChatResponse),
		Chats: http.NewServer(en.Chats, decodeHTTPChatsRequest, encodeHTTPChatResponse),
	}
	return handler
}
