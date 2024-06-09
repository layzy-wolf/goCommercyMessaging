package service

import (
	pb "GroupService/api/group.v1"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Group struct {
	Name  string   `bson:"name"`
	User  string   `bson:"user"`
	Users []string `bson:"users"`
	Code  string   `bson:"code"`
}

type User struct {
	Login    string   `bson:"login"`
	Password string   `bson:"password"`
	Token    string   `bson:"token"`
	Chats    []string `bson:"chats"`
}

func (s *Service) AddToGroup(ctx context.Context, in *pb.AddReq) (*pb.BoolResp, error) {
	var gr Group
	var usr User

	col := s.client.Database("group").Collection("groups")
	token, _ := jwt.Parse(in.Code, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &secret, nil
	})

	if err := col.FindOne(ctx, bson.M{"name": fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["name"])}).Decode(&gr); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	if gr.Code != in.Code {
		return nil, errors.New("invalid code")
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

	col = s.client.Database("auth").Collection("users")

	if err := col.FindOne(ctx, bson.M{"login": in.User}).Decode(&usr); err != nil {
		log.Printf("%v", err)
	} else {
		usr.Chats = append(usr.Chats, gr.Name)
		if _, err := col.UpdateOne(ctx, bson.M{"login": usr.Login}, bson.M{"$set": usr}); err != nil {
			log.Printf("%v", err)
		}
	}

	return &pb.BoolResp{Success: true}, nil
}

func (s *Service) RemoveFromGroup(ctx context.Context, in *pb.RemoveReq) (*pb.BoolResp, error) {
	var (
		gr    Group
		nu    []string
		usr   User
		chats []string
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

	col = s.client.Database("auth").Collection("users")

	if err := col.FindOne(ctx, bson.M{"login": in.User}).Decode(&usr); err != nil {
		log.Printf("%v", err)
	} else {
		for _, v := range usr.Chats {
			if v != gr.Name {
				chats = append(chats, v)
			}
		}
		usr.Chats = chats

		if _, err := col.UpdateOne(ctx, bson.M{"login": usr.Login}, bson.M{"$set": usr}); err != nil {
			log.Printf("%v", err)
		}
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
