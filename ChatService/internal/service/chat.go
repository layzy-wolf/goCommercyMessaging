package service

import (
	pb "ChatService/api/grpc/chat.v1"
	store "ChatService/api/grpc/chatStore.v1"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (s *Server) GetMessages(ctx context.Context, in *pb.GetRequest) (*pb.ChatMessages, error) {
	var msg []*pb.Message

	res, err := client.GetMessages(ctx, &store.GetReq{
		From:      in.From,
		To:        in.To,
		Limit:     in.Limit,
		Timestamp: in.Timestamp,
	})

	if err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	for _, r := range res.Messages {
		msgHC := jwt.NewWithClaims(jwt.SigningMethodHS256, MessageHead{
			From: r.From,
			To:   r.To,
		})
		msgHT, _ := msgHC.SignedString(Secret)

		msgBC := jwt.NewWithClaims(jwt.SigningMethodHS256, MessageBody{
			Message:   r.Message,
			Timestamp: r.Timestamp,
		})
		msgBT, _ := msgBC.SignedString(Secret)

		msg = append(msg, &pb.Message{
			MessageHead: msgHT,
			MessageBody: msgBT,
		})
	}

	return &pb.ChatMessages{Message: msg}, nil
}

func (s *Server) ForwardMessage(stream pb.Chat_ForwardMessageServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	done = make(chan bool)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	md, _ := metadata.FromIncomingContext(stream.Context())
	u := getUserFromMD(md)

	if err := register(u, &stream); err != nil {
		log.Printf("E: %v", err)
		return err
	}

	go listen(ctx, stream)
	go forward(ctx)

	select {
	case <-ctx.Done():
		return nil
	}
}

func register(u string, stream *pb.Chat_ForwardMessageServer) error {
	if u == "" {
		return errors.New("user key is undefined")
	}
	list.Insert(u, stream)
	return nil
}

func forget(u string) error {
	err := list.Delete(u)
	if err != nil {
		return err
	}
	return nil
}

func getUserFromMD(md metadata.MD) (user string) {
	usr, _ := jwt.Parse(md["jwtuser"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})
	return fmt.Sprintf("%v", usr.Claims.(jwt.MapClaims)["user"])
}
