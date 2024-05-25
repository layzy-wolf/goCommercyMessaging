package service

import (
	pb "GroupService/api/group.v1"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (s *Service) Register(ctx context.Context, in *pb.RegisterReq) (*pb.BoolResp, error) {
	var g Group

	col := s.client.Database("group").Collection("groups")

	cond := bson.M{
		"name": "#" + in.Name,
	}

	err := col.FindOne(ctx, cond).Decode(&g)
	if err == nil {
		return &pb.BoolResp{Success: false}, errors.New("group is exist")
	}

	d := Group{
		Name:  "#" + in.Name,
		User:  in.User,
		Users: []string{in.User},
		Code:  "",
	}

	if _, err := col.InsertOne(ctx, d); err != nil {
		log.Printf("E: %v", err)
		return &pb.BoolResp{Success: false}, err
	}
	return &pb.BoolResp{Success: true}, nil
}
