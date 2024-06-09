package test

import (
	pb "ChatService/api/grpc/chat.v1"
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"testing"
)

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

func TestChat(t *testing.T) {

	/*
		Тестирование сервиса чата
		Для начала подключаемся к сервису чата и
		передаем в него тестовую учетную запись,
		после шифруем и отправляем сообщение,
		если подтверждение сообщения отправляется обратно,
		то тест считается выполненным
	*/

	login := "test1"
	secret := []byte("my-Secret")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	assert.Nil(t, err)

	c := pb.NewChatClient(cc)

	j := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTUser{
		login,
		jwt.StandardClaims{
			ExpiresAt: 1500,
			Issuer:    "issuer",
		},
	})

	ss, err := j.SignedString(secret)

	assert.Nil(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "jwtuser", ss)

	stream, err := c.ForwardMessage(ctx)

	assert.Nil(t, err)

	msgH := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTMsgHead{
		From:           login,
		To:             "test",
		StandardClaims: jwt.StandardClaims{},
	})

	msgB := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTMsgBody{
		Message:        "test",
		Timestamp:      "",
		StandardClaims: jwt.StandardClaims{},
	})

	msgHS, err := msgH.SignedString(secret)

	assert.Nil(t, err)

	msgBS, err := msgB.SignedString(secret)

	assert.Nil(t, err)

	err = stream.Send(&pb.Message{
		MessageHead: msgHS,
		MessageBody: msgBS,
	})

	assert.Nil(t, err)

	msg, err := stream.Recv()

	assert.Nil(t, err)
	assert.NotEmpty(t, msg)
}
