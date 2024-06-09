package test

import (
	pb "AuthService/api/auth.v1"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestAuth(t *testing.T) {

	/*
		Тестирование сервиса автоизации
		Для начала берем тестовую учетную запись
		и входим в систему, после получаем код верификации
		и обращаемся с запосом верификации,
		если верификация прошла успешно, то тест пройден успешно.
	*/

	login := "test1"
	password := "test1"

	cc, err := grpc.Dial("localhost:40444", grpc.WithTransportCredentials(insecure.NewCredentials()))

	assert.Nil(t, err)

	c := pb.NewAuthClient(cc)

	resp, err := c.Login(context.TODO(), &pb.UserRequest{
		Login:  login,
		Passwd: password,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, resp.Token)

	res, err := c.Verify(context.TODO(), &pb.VerifyReq{Token: resp.Token})

	assert.Nil(t, err)
	assert.Equal(t, true, res.Success)
}
