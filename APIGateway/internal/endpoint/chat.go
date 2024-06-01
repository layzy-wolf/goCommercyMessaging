package endpoint

import (
	"ApiGateway/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
)

type ChatEndpoints struct {
	Show  endpoint.Endpoint
	Chats endpoint.Endpoint
}

type ShowRequest struct {
	From  string
	To    string `json:"to"`
	Limit int64  `json:"limit"`
}

type ChatsRequest struct {
	User      string
	Timestamp string `json:"timestamp"`
}

type ShowMessagesResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	MessageHead string `json:"messageHead"`
	MessageBody string `json:"messageBody"`
}

func MakeChatEndpoints(srv service.ChatService) ChatEndpoints {
	return ChatEndpoints{
		Show:  MakeShowEndpoint(srv),
		Chats: MakeChatsEndpoint(srv),
	}
}

func MakeShowEndpoint(srv service.ChatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var mes ShowMessagesResponse

		req := request.(ShowRequest)
		messages, err := srv.GetMessages(req.From, req.To, req.Limit)
		if err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		for _, v := range messages {
			mes.Messages = append(mes.Messages, Message{
				MessageHead: v["MessageHead"],
				MessageBody: v["MessageBody"],
			})
		}

		return mes, nil
	}
}

func MakeChatsEndpoint(srv service.ChatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var mes ShowMessagesResponse

		req := request.(ChatsRequest)
		messages, err := srv.GetChats(req.User, req.Timestamp)
		if err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		for _, v := range messages {
			mes.Messages = append(mes.Messages, Message{
				MessageHead: v["MessageHead"],
				MessageBody: v["MessageBody"],
			})
		}

		return mes, nil
	}
}
