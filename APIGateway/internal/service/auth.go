package service

import (
	"ApiGateway/config"
	pb "AuthService/api/auth.v1"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type authService struct {
	client pb.AuthClient
}

func NewAuthService(cfg config.Cfg) AuthService {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", cfg.Auth.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("E: %v", err)
	}
	log.Printf("Auth service: %v", cc.GetState())
	c := pb.NewAuthClient(cc)
	return &authService{client: c}
}

func (s *authService) Register(token string) (bool, error) {
	j, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})

	resp, err := s.client.Register(context.Background(), &pb.UserRequest{
		Login:  fmt.Sprintf("%v", j.Claims.(jwt.MapClaims)["login"]),
		Passwd: fmt.Sprintf("%v", j.Claims.(jwt.MapClaims)["password"]),
	})
	if err != nil {
		log.Printf("%v", err)
		return false, err
	}

	return resp.Success, err
}

func (s *authService) Login(token string) (string, error) {
	j, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})

	resp, err := s.client.Login(context.Background(), &pb.UserRequest{
		Login:  fmt.Sprintf("%v", j.Claims.(jwt.MapClaims)["login"]),
		Passwd: fmt.Sprintf("%v", j.Claims.(jwt.MapClaims)["password"]),
	})
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	return resp.Token, err
}

func (s *authService) Verify(token string) bool {
	resp, _ := s.client.Verify(context.Background(), &pb.VerifyReq{Token: token})
	return resp.Success
}
