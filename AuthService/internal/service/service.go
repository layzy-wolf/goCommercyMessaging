package service

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	pb "github.com/layzy-wolf/goCommercyMessaging/AuthService/api/auth"
	"github.com/layzy-wolf/goCommercyMessaging/AuthService/internal/storage"
	"time"
)

var (
	simpleKey = "simpleKey"
)

type Server struct {
	storage *storage.Store
	pb.UnimplementedAuthServer
}

func New(storage *storage.Store) *Server {
	return &Server{
		storage:                 storage,
		UnimplementedAuthServer: pb.UnimplementedAuthServer{},
	}
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := storage.User{
		Login:    in.Login,
		Password: in.Passwd,
		Email:    in.Email,
		TTLToken: time.Now(),
	}

	res := s.storage.DB.Create(user)

	return &pb.RegisterResponse{UserId: int32(user.ID)}, res.Error
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := s.storage.DB.Take(storage.User{}).Where(storage.User{
		Email:    in.Email,
		Password: in.Passwd,
	})

	if user.Error != nil {
		return &pb.LoginResponse{Token: "your credentials don`t match our records"}, user.Error
	}

	s.storage.DB.Model(&user).Update("TTLToken", time.Now().AddDate(0, 1, 0))

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": s.storage.DB.Model(&user).Take("Login"),
	})

	token, _ := tkn.SignedString([]byte(simpleKey))

	return &pb.LoginResponse{Token: token}, nil
}

func (s *Server) VerifyAction(ctx context.Context, in *pb.VerifyToken) (*pb.VerifyResponse, error) {
	token, _ := jwt.Parse(&in.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		return []byte(simpleKey), nil
	})
	//TODO
	if claims, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		fmt.Println(claims["user"])
	}
	return &pb.VerifyResponse{}, nil
}
