package service

import (
	"ApiGateway/config"
	pb "GroupService/api/group.v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type groupService struct {
	client pb.GroupClient
}

func NewGroupChatService(cfg config.Cfg) GroupService {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", cfg.Group.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("E: %v", err)
	}
	log.Printf("Group service: %v", cc.GetState())
	c := pb.NewGroupClient(cc)
	return &groupService{client: c}
}

func (s *groupService) Register(user, name string) bool {
	resp, err := s.client.Register(context.Background(), &pb.RegisterReq{
		User: user,
		Name: name,
	})
	if err != nil {
		log.Printf("%v", err)
		return false
	}
	return resp.Success
}

func (s *groupService) GetCode(name, user string) (string, error) {
	resp, err := s.client.GetCode(context.Background(), &pb.GetReq{
		Name: name,
		User: user,
	})
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	return resp.Code, nil
}

func (s *groupService) AddToGroup(token, user string) bool {
	resp, err := s.client.AddToGroup(context.Background(), &pb.AddReq{
		Code: token,
		User: user,
	})
	if err != nil {
		log.Printf("%v", err)
		return false
	}
	return resp.Success
}

func (s *groupService) RemoveFromGroup(user, group, remove string) bool {
	resp, err := s.client.RemoveFromGroup(context.Background(), &pb.RemoveReq{
		User:   user,
		Group:  group,
		Remove: remove,
	})
	if err != nil {
		log.Printf("%v", err)
		return false
	}
	return resp.Success
}

func (s *groupService) GetMembers(name, user string) []string {
	resp, err := s.client.GetMembers(context.Background(), &pb.GetReq{
		Name: name,
		User: user,
	})
	if err != nil {
		log.Printf("%v", err)
		return nil
	}
	return resp.Members
}
