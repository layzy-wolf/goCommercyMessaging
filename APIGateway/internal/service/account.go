package service

import (
	pb "AccountService/api/account.v1"
	"ApiGateway/config"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type accountService struct {
	client pb.AccountClient
}

func NewAccountService(cfg *config.Cfg) *accountService {
	cc, err := grpc.Dial(fmt.Sprintf("%v", cfg.Account.Host), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("E: %v", err)
	}
	log.Printf("Account service: %v", cc.GetState())
	c := pb.NewAccountClient(cc)
	return &accountService{client: c}
}

func (s *accountService) Register(token string) (bool, error) {
	j, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})

	resp, err := s.client.Register(context.Background(), &pb.RegisterRequest{
		Login:    fmt.Sprintf("%v", j.Claims.(jwt.MapClaims)["login"]),
		Password: fmt.Sprintf("%v", j.Claims.(jwt.MapClaims)["password"]),
	})

	if err != nil {
		log.Printf("%v", err)
		return false, err
	}

	return resp.Success, err
}

func (s *accountService) Remove(user string) bool {
	resp, err := s.client.Remove(context.Background(), &pb.UserRequest{
		Login: fmt.Sprintf("%v", user),
	})

	if err != nil {
		log.Printf("%v", err)
	}

	return resp.Success
}

func (s *accountService) List(token string) ([]string, error) {
	j, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})

	resp, err := s.client.List(context.Background(), &pb.UserRequest{
		Login: fmt.Sprintf("%v", j.Claims.(jwt.MapClaims)["login"]),
	})

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return resp.User, err
}

func (s *accountService) Search(condition string) ([]string, error) {
	resp, err := s.client.Search(context.Background(), &pb.SearchRequest{
		Condition: condition,
	})

	if err != nil {
		return nil, err
	}

	return resp.User, err
}

func (s *accountService) AddContact(user, chat string) (bool, error) {
	resp, err := s.client.AddContact(context.Background(), &pb.ContactRequest{
		User: user,
		Chat: chat,
	})

	if err != nil {
		log.Printf("%v", err)
		return false, err
	}

	return resp.Success, err
}

func (s *accountService) RemoveContact(user, chat string) (bool, error) {
	resp, err := s.client.RemoveContact(context.Background(), &pb.ContactRequest{
		User: user,
		Chat: chat,
	})

	if err != nil {
		log.Printf("%v", err)
	}

	return resp.Success, err
}
