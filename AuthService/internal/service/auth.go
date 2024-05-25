package service

import (
	pb "AuthService/api/auth.v1"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	Login    string `bson:"login"`
	Password string `bson:"password"`
	Token    string `bson:"token"`
}

func (s *Service) Register(ctx context.Context, in *pb.UserRequest) (*pb.BoolResp, error) {
	var usr User

	passwd, _ := hash(in.Passwd)

	cond := bson.M{
		"login": "@" + in.Login,
	}

	u := User{
		Login:    "@" + in.Login,
		Password: passwd,
		Token:    "",
	}

	col := getCol()
	err := col.FindOne(ctx, cond).Decode(&usr)
	if err == nil {
		return &pb.BoolResp{Success: false}, errors.New("user is exist")
	}

	if _, err := col.InsertOne(ctx, u); err != nil {
		log.Printf("%v", err)
		return &pb.BoolResp{Success: false}, err
	}
	return &pb.BoolResp{Success: true}, nil
}

func (s *Service) Login(ctx context.Context, in *pb.UserRequest) (*pb.LoginResp, error) {
	var usr User

	cond := bson.M{
		"login": in.Login,
	}

	col := getCol()
	err := col.FindOne(ctx, cond).Decode(&usr)

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	if f := verify(usr.Password, in.Passwd); f {
		jwtTkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"login": usr.Login,
			"time":  time.Now(),
		})
		signedJwt, err := jwtTkn.SignedString(Secret)
		if err != nil {
			log.Printf("%v", err)
		}

		usr.Token = signedJwt

		if _, err := col.UpdateOne(ctx, bson.M{"login": usr.Login}, bson.M{
			"$set": usr,
		}); err != nil {
			log.Printf("E: %v", err)
			return nil, err
		}

		return &pb.LoginResp{Token: signedJwt}, nil
	}

	return nil, errors.New("user data don't math our records")
}

func (s *Service) Verify(ctx context.Context, in *pb.VerifyReq) (*pb.VerifyResp, error) {
	var user User

	token, _ := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})

	col := getCol()
	login := fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["login"])

	if err := col.FindOne(ctx, bson.M{"login": login}).Decode(&user); err != nil {
		log.Printf("E: %v", err)
		return nil, err
	}

	if user.Token == in.Token {
		return &pb.VerifyResp{
			Success: true,
			User:    login,
		}, nil
	}

	return nil, errors.New("token don't match our records")
}

func getCol() *mongo.Collection {
	return c.Database("auth").Collection("users")
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func verify(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
