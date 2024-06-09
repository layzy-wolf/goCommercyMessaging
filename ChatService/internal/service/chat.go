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

func (s *Server) Test(_ context.Context, in *pb.Bool) (*pb.Bool, error) {
	log.Println("called")
	return &pb.Bool{Success: true}, nil
}

func (s *Server) GetMessages(ctx context.Context, in *pb.GetRequest) (*pb.ChatMessages, error) {

	// Входные параметры: отправить, получатель, лимит и время

	/*
		Принцип работы: обраится к сервису-хранилищу и
		либо зашифровать все сообщения,
		либо вернуть ошибку
	*/

	// Выходные данные: массив сообщений и ошибка

	var msg []*pb.Message
	var err error
	res := &store.ChatMessages{}

	if in.Timestamp == "" {
		res, err = client.GetMessages(ctx, &store.GetRequest{
			From:  in.From,
			To:    in.To,
			Limit: in.Limit,
		})
	} else {
		res, err = client.UpdateChats(ctx, &store.UpdateRequest{
			From:      in.From,
			To:        in.To,
			Timestamp: in.Timestamp,
		})
	}

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

	// Входные параметры: stream пользователя

	/*
		Принцип работы: при вызове получить пользователя из метаинформации
		и зарегестрировать его во внутреннем связанном списке в системе,
		после прослушивать его канал связи.
		func listen - работает для прослушивания,
		горутина прерывается до получения сообщения,
		после прерывание спадает, сообщение передается в forward.
		func forward - работает для передачи,
		горутина прерывается до получения сообщений,
		полсе получения сообщения прерывание снимается,
		в связанном списке находит адресата и отправляет ему сообщение
		и снова прерывается.
	*/

	// Выходные данные: ошибка

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
