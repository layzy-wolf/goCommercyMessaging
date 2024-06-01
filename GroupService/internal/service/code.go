package service

import (
	pb "GroupService/api/group.v1"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Code struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	jwt.StandardClaims
}

func (s *Service) GetCode(ctx context.Context, in *pb.GetReq) (*pb.GetCodeResp, error) {
	var gr Group

	col := s.client.Database("group").Collection("groups")

	if err := col.FindOne(ctx, bson.M{"name": in.Name}).Decode(&gr); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	if in.User != gr.User {
		return nil, errors.New("forbidden")
	}

	if gr.Code == "" {
		jwtString := jwt.NewWithClaims(jwt.SigningMethodHS256, Code{
			Name:      gr.Name,
			Timestamp: fmt.Sprintf("%v", time.Now()),
		})
		gr.Code, _ = jwtString.SignedString(secret)
		if _, err := col.UpdateOne(ctx, bson.M{"name": gr.Name}, bson.M{
			"$set": gr,
		}); err != nil {
			log.Printf("E: %v", err)
			return nil, err
		}
	}

	return &pb.GetCodeResp{Code: gr.Code}, nil
}
