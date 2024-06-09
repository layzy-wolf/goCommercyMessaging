package store

import (
	pb "ChatService/api/grpc/chatStore.v1"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Message struct {
	From      string    `bson:"from"`
	Message   string    `bson:"message"`
	To        string    `bson:"to"`
	Timestamp time.Time `bson:"timestamp"`
}

type Group struct {
	Name  string   `bson:"name"`
	User  string   `bson:"user"`
	Users []string `bson:"users"`
	Code  string   `bson:"code"`
}

func (s *Server) Test(_ context.Context, in *pb.Bool) (*pb.Bool, error) {
	log.Println("called")
	return &pb.Bool{Success: true}, nil
}

func (s *Server) GetMessages(ctx context.Context, in *pb.GetRequest) (*pb.ChatMessages, error) {
	var (
		res  []Message
		msg  []*pb.ChatMessage
		cond bson.M
		opts options.FindOptions
	)

	col := getCol()

	opts.SetSort(bson.M{"timestamp": -1})
	if in.To[0:1] == "&" {
		cond = bson.M{
			"to": in.To,
		}
	} else {
		cond = bson.M{
			"$or": bson.A{
				bson.M{
					"to":   in.From,
					"from": in.To,
				},
				bson.M{
					"from": in.From,
					"to":   in.To,
				},
			},
		}
	}

	opts.SetLimit(in.Limit)

	cur, err := col.Find(ctx, cond, &opts)
	if err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	if err = cur.All(ctx, &res); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	for _, r := range res {
		msg = append(msg, &pb.ChatMessage{
			From:      r.From,
			Message:   r.Message,
			To:        r.To,
			Timestamp: fmt.Sprintf("%v", r.Timestamp),
		})
	}
	return &pb.ChatMessages{Messages: msg}, nil
}

func (s *Server) UpdateChats(ctx context.Context, in *pb.UpdateRequest) (*pb.ChatMessages, error) {
	var (
		gr  []Group
		res []Message
		msg []*pb.ChatMessage
	)

	timestamp, _ := time.Parse(time.DateTime, in.Timestamp)

	grCol := c.Database("group").Collection("groups")

	grCond := bson.M{
		"users": bson.M{
			"$in": bson.A{in.From},
		},
	}

	grCur, _ := grCol.Find(ctx, grCond)

	grCur.All(ctx, &gr)

	condArr := bson.A{
		bson.M{
			"to": in.From,
		},
		bson.M{
			"from": in.From,
		},
	}

	for _, g := range gr {
		condArr = append(condArr, bson.M{"to": g.Name})
	}

	cond := bson.M{
		"$or": condArr,
		"timestamp": bson.M{
			"$gte": timestamp,
		},
	}

	col := getCol()

	cur, err := col.Find(ctx, cond)
	if err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	if err = cur.All(ctx, &res); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	for _, r := range res {
		msg = append(msg, &pb.ChatMessage{
			From:      r.From,
			Message:   r.Message,
			To:        r.To,
			Timestamp: fmt.Sprintf("%v", r.Timestamp),
		})
	}
	return &pb.ChatMessages{Messages: msg}, nil
}

func (s *Server) AddMessage(ctx context.Context, in *pb.ChatMessage) (*pb.BoolResp, error) {
	col := getCol()
	_, err := col.InsertOne(ctx, bson.M{
		"from":      in.From,
		"to":        in.To,
		"message":   in.Message,
		"timestamp": time.Now().UTC(),
	})
	if err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	return &pb.BoolResp{Successful: true}, nil
}

func (s *Server) EditMessage(ctx context.Context, in *pb.ChatMessage) (*pb.BoolResp, error) {
	t, _ := time.Parse(time.RFC3339, in.Timestamp)
	msg := Message{From: in.From, Message: in.Message, To: in.To, Timestamp: t}

	filter := bson.M{
		"from":      msg.From,
		"timestamp": msg.Timestamp,
	}

	col := getCol()
	_, err := col.UpdateOne(ctx, filter, bson.M{"$set": msg})
	if err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	return &pb.BoolResp{Successful: true}, nil

}

func (s *Server) GetUsersFromGroup(ctx context.Context, in *pb.Group) (*pb.Members, error) {
	var usrs []string
	var res Group

	col := c.Database("group").Collection("groups")

	cursor := col.FindOne(ctx, bson.M{"name": in.Name})

	if err := cursor.Decode(&res); err != nil {
		log.Printf("%v", err)
	}

	for _, r := range res.Users {
		usrs = append(usrs, fmt.Sprintf("%v", r))
	}
	return &pb.Members{Name: usrs}, nil
}

func getCol() *mongo.Collection {
	col := c.Database("chat").Collection("messages")
	return col
}
