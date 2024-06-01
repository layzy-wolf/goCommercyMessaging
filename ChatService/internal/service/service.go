package service

import (
	pb "ChatService/api/grpc/chat.v1"
	store "ChatService/api/grpc/chatStore.v1"
	"github.com/golang-jwt/jwt"
)

var (
	list   LinkedList
	done   chan bool
	msg    pb.Message
	client store.StoreClient
	Secret = []byte("my-Secret")
)

type MessageHead struct {
	From string `json:"from"`
	To   string `json:"to"`
	jwt.StandardClaims
}

type MessageBody struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	jwt.StandardClaims
}

type Server struct {
	Message map[string]*pb.Message
	pb.UnimplementedChatServer
}

func New(c store.StoreClient) *Server {
	client = c
	s := &Server{Message: make(map[string]*pb.Message)}
	return s
}
