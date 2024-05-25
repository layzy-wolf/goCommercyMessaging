package test

import (
	"context"
	pb "github.com/layzy-wolf/goCommercyMessaging/AuthService/api/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/randomstring.v1"
	"testing"
	"time"
)

var (
	addr   = "localhost:40444"
	email  = randomstring.Generate(8, randomstring.LowerLetters, "@a")
	login  = randomstring.Generate(8, randomstring.LowerLetters)
	passwd = randomstring.Generate(8, randomstring.LowerLetters)
)

func connectToGrpc() (connection *grpc.ClientConn, service pb.AuthClient, err error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}

	c := pb.NewAuthClient(conn)

	return conn, c,nil
}

func TestRegsister(t *testing.T) {
	conn, c, err := connectToGrpc()
	if err != nil {
		t.Fatalf("error while connection: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := c.Register(ctx, &pb.RegisterRequest{
		Email:  email,
		Login:  login,
		Passwd: passwd,
	})

	if err != nil {
		t.Fatalf("register error: %v", err)
	}

	t.Logf("user id: %v", res.UserId)
}

func TestLogin(t *testing.T) {
	conn, c, err := connectToGrpc()
	if err != nil {
		t.Fatalf("error while connection: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := c.Login(ctx, &pb.LoginRequest{
		Email: email,
		Passwd: passwd,
	})

	if err != nil {
		t.Fatalf("login error: %v", err)
	}

	t.Logf("user token: %v", res.Token)
}