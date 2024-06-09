package service

import (
	pb "AccountService/api/account.v1"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.BoolResponse, error) {
	// Входные параметры: Логин и пароль пользователя

	/*
		Принцип работы: Найти пользователя по логину, если такого нет,
		то внести пользователя в БД
	*/

	// Выходные данные: булевое значение и ошибка

	var usr User

	passwd, _ := hash(in.Password)

	cond := bson.M{
		"login": "@" + in.Login,
	}

	u := User{
		Login:    "@" + in.Login,
		Password: passwd,
		Token:    "",
		Chats:    nil,
	}

	col := s.conn.Database("auth").Collection("users")
	err := col.FindOne(ctx, cond).Decode(&usr)
	if err == nil {
		return &pb.BoolResponse{Success: false}, errors.New("user is exist")
	}

	if _, err := col.InsertOne(ctx, u); err != nil {
		log.Printf("%v", err)
		return &pb.BoolResponse{Success: false}, err
	}
	return &pb.BoolResponse{Success: true}, nil
}

func (s *Service) Remove(ctx context.Context, in *pb.UserRequest) (*pb.BoolResponse, error) {
	cond := bson.M{
		"login": in.Login,
	}

	col := s.conn.Database("auth").Collection("users")

	if _, err := col.DeleteOne(ctx, cond); err != nil {
		return &pb.BoolResponse{Success: false}, err
	}

	return &pb.BoolResponse{Success: true}, nil
}

func (s *Service) List(ctx context.Context, in *pb.UserRequest) (*pb.UsersList, error) {
	var user User

	cond := bson.M{
		"login": in.Login,
	}

	col := s.conn.Database("auth").Collection("users")

	if err := col.FindOne(ctx, cond).Decode(&user); err != nil {
		return nil, err
	}

	return &pb.UsersList{User: user.Chats}, nil
}

func (s *Service) Search(ctx context.Context, in *pb.SearchRequest) (*pb.UsersList, error) {
	var users []User
	var usersList []string

	cond := bson.M{
		"login": bson.M{
			"$regex": in.Condition,
		},
	}

	col := s.conn.Database("auth").Collection("users")

	res, err := col.Find(ctx, cond)
	if err != nil {
		return nil, err
	}
	if err = res.All(ctx, &users); err != nil {
		return nil, err
	}

	for _, u := range users {
		usersList = append(usersList, u.Login)
	}

	return &pb.UsersList{User: usersList}, err
}

func (s *Service) AddContact(ctx context.Context, in *pb.ContactRequest) (*pb.BoolResponse, error) {
	var usr User

	cond := bson.M{"login": in.User}

	col := s.conn.Database("auth").Collection("users")

	if err := col.FindOne(ctx, cond).Decode(&usr); err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	usr.Chats = append(usr.Chats, in.Chat)

	if _, err := col.UpdateOne(ctx, cond, bson.M{"$set": usr}); err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return &pb.BoolResponse{Success: true}, nil
}

func (s *Service) RemoveContact(ctx context.Context, in *pb.ContactRequest) (*pb.BoolResponse, error) {
	var (
		usr   User
		chats []string
	)

	cond := bson.M{"login": in.User}

	col := s.conn.Database("auth").Collection("users")

	if err := col.FindOne(ctx, cond).Decode(&usr); err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	for _, v := range usr.Chats {
		if v != in.Chat {
			chats = append(chats, v)
		}
	}

	usr.Chats = chats

	if _, err := col.UpdateOne(ctx, cond, bson.M{"$set": usr}); err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	if in.Chat[0:1] == "&" {
		col = s.conn.Database("group").Collection("groups")
		var gr Group
		var grs []string

		col.FindOne(ctx, bson.M{"name": in.Chat}).Decode(&gr)

		for _, u := range gr.Users {
			if u != in.User {
				grs = append(grs, u)
			}
		}

		gr.Users = grs

		if _, err := col.UpdateOne(ctx, bson.M{"name": gr.Name}, bson.M{"$set": gr}); err != nil {
			return &pb.BoolResponse{Success: false}, err
		}
	}

	return &pb.BoolResponse{Success: true}, nil
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
