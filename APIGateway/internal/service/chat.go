package service

import (
	"ApiGateway/config"
	pb "ChatService/api/grpc/chat.v1"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

type chatService struct {
	client pb.ChatClient
}

type JWTMessage struct {
	MessageHead string `json:"MessageHead"`
	MessageBody string `json:"MessageBody"`
	jwt.StandardClaims
}

type JWTUser struct {
	User string `json:"user"`
	jwt.StandardClaims
}

type JWTMsgHead struct {
	From string `json:"from"`
	To   string `json:"to"`
	jwt.StandardClaims
}

type JWTMsgBody struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	jwt.StandardClaims
}

func NewChatService(cfg config.Cfg) ChatService {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", cfg.Chat.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("E: %v", err)
	}
	log.Printf("Chat service: %v", cc.GetState())
	c := pb.NewChatClient(cc)
	return &chatService{client: c}
}

func (s *chatService) GetChats(user, time string) ([]map[string]string, error) {
	var message []map[string]string
	res, err := s.client.GetMessages(context.Background(), &pb.GetRequest{
		From:      user,
		To:        "",
		Limit:     0,
		Timestamp: time,
	})
	for _, m := range res.Message {
		message = append(message, map[string]string{"MessageHead": m.MessageHead, "MessageBody": m.MessageBody})
	}
	return message, err
}

func (s *chatService) GetMessages(from, to string, limit int64) ([]map[string]string, error) {
	var message []map[string]string
	res, err := s.client.GetMessages(context.Background(), &pb.GetRequest{
		From:      from,
		To:        to,
		Limit:     limit,
		Timestamp: "",
	})
	for _, m := range res.Message {
		message = append(message, map[string]string{"MessageHead": m.MessageHead, "MessageBody": m.MessageBody})
	}
	return message, err
}

func (s *chatService) ForwardMessage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("%v", err)
		cancel()
	}

	j := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTUser{
		r.Header.Get("user"),
		jwt.StandardClaims{
			ExpiresAt: 1500,
			Issuer:    "issuer",
		},
	})
	ss, _ := j.SignedString(Secret)
	ctx = metadata.AppendToOutgoingContext(ctx, "jwtuser", ss)

	cl, err := s.client.ForwardMessage(ctx)
	if err != nil {
		log.Printf("%v", err)
		cancel()
	}

	go read(ctx, cancel, conn, cl)
	go write(ctx, cancel, conn, cl)

	select {
	case <-ctx.Done():
		return
	}
}

func read(ctx context.Context, cancel context.CancelFunc, conn *websocket.Conn, cl pb.Chat_ForwardMessageClient) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("%v", err)
			cancel()
			return
		}
		j, _ := jwt.Parse(string(p), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return &Secret, nil
		})
		data := j.Claims.(jwt.MapClaims)

		jH := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTMsgHead{
			data["head"].(map[string]interface{})["from"].(string),
			data["head"].(map[string]interface{})["to"].(string),
			jwt.StandardClaims{
				ExpiresAt: 1500,
				Issuer:    "issuer",
			},
		})

		jB := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTMsgBody{
			data["body"].(map[string]interface{})["message"].(string),
			fmt.Sprintf("%v", data["body"].(map[string]interface{})["timestamp"]),
			jwt.StandardClaims{
				ExpiresAt: 1500,
				Issuer:    "issuer",
			},
		})

		signedJH, _ := jH.SignedString(Secret)
		signedJB, _ := jB.SignedString(Secret)

		if err := cl.Send(&pb.Message{
			MessageHead: signedJH,
			MessageBody: signedJB,
		}); err != nil {
			log.Printf("%v", err)
			return
		}
		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}

func write(ctx context.Context, cancel context.CancelFunc, conn *websocket.Conn, cl pb.Chat_ForwardMessageClient) {
	for {
		in, err := cl.Recv()
		if err != nil {
			log.Printf("%v", err)
			cancel()
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTMessage{
			in.MessageHead,
			in.MessageBody,
			jwt.StandardClaims{
				ExpiresAt: 1500,
				Issuer:    "issuer",
			},
		})
		s, err := token.SignedString(Secret)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		if err := conn.WriteMessage(websocket.TextMessage, []byte(s)); err != nil {
			log.Printf("%v", err)
			return
		}
		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}
