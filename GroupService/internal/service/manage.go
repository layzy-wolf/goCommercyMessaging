package service

import (
	pb "GroupService/api/group.v1"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Group struct {
	Name  string   `bson:"name"`
	User  string   `bson:"user"`
	Users []string `bson:"users"`
	Code  string   `bson:"code"`
}

func (s *Service) AddToGroup(ctx context.Context, in *pb.AddReq) (*pb.BoolResp, error) {
	var gr Group

	col := s.client.Database("group").Collection("groups")

	if err := col.FindOne(ctx, bson.M{"code": in.Code}).Decode(&gr); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	for _, u := range gr.Users {
		if u == in.User {
			return nil, errors.New("user already joined")
		}
	}
	gr.Users = append(gr.Users, in.User)

	if _, err := col.UpdateOne(ctx, bson.M{"name": gr.Name}, bson.M{
		"$set": gr,
	}); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}
	return &pb.BoolResp{Success: true}, nil
}

func (s *Service) RemoveFromGroup(ctx context.Context, in *pb.RemoveReq) (*pb.BoolResp, error) {
	var (
		gr Group
		nu []string
	)

	col := s.client.Database("group").Collection("groups")

	if err := col.FindOne(ctx, bson.M{"name": in.Group}).Decode(&gr); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	if gr.User != in.User {
		return nil, errors.New("you don't have permission")
	}

	for _, u := range gr.Users {
		if u != in.Remove {
			nu = append(nu, u)
		}
	}

	gr.Users = nu

	if _, err := col.UpdateOne(ctx, bson.M{"name": gr.Name}, bson.M{
		"$set": gr,
	}); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}
	return &pb.BoolResp{Success: true}, nil
}

func (s *Service) GetMembers(ctx context.Context, in *pb.GetReq) (*pb.GetMembersResp, error) {
	var gr Group

	col := s.client.Database("group").Collection("groups")

	if err := col.FindOne(ctx, bson.M{"name": in.Name}).Decode(&gr); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	suc := false

	for _, v := range gr.Users {
		if v == in.User {
			suc = true
		}
	}

	if !suc {
		return nil, errors.New("forbidden")
	}

	return &pb.GetMembersResp{Members: gr.Users}, nil
}
