package transport

import (
	"ApiGateway/config"
	pb "AuthService/api/auth.v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	net "net/http"
)

type V interface {
	Verify(r net.Request) (bool, error)
}

type Verify struct {
	client pb.AuthClient
}

func NewVerify(cfg config.Cfg) *Verify {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", cfg.Auth.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("E: %v", err)
	}
	c := pb.NewAuthClient(cc)
	return &Verify{
		client: c,
	}
}

func (v *Verify) Verify(r *net.Request) (bool, string, error) {
	auth := fmt.Sprintf("%v", r.Header.Get("Authorization"))
	resp, err := v.client.Verify(context.Background(), &pb.VerifyReq{Token: auth})
	if err != nil {
		log.Printf("%v", err)
		return false, "", err
	}
	return resp.Success, resp.User, nil
}
