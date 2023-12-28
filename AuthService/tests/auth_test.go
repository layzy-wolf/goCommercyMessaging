package tests

import (
	authv1 "app/api/protobuf/auth.v1"
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

const (
	appSecret = "test_secret"

	passDefaultLen = 10
)

func TestAuthLoginHappyPath(t *testing.T) {
	ctx := context.Background()

	grpcAdd := net.JoinHostPort("localhost", "40444")
	cc, _ := grpc.DialContext(
		ctx,
		grpcAdd,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	client := authv1.NewAuthClient(cc)

	email := gofakeit.Email()
	pass := gofakeit.Password(true, true, true, true, false, passDefaultLen)

	respReg, err := client.Register(ctx, &authv1.RegisterRequest{
		Email:  email,
		Passwd: pass,
	})

	require.NoError(t, err)
	require.NotEmpty(t, respReg.GetUserId())

	respLogin, err := client.Login(ctx, &authv1.LoginRequest{
		Email:  email,
		Passwd: pass,
	})

	require.NoError(t, err)

	token := respLogin.GetToken()
	require.NotEmpty(t, token)

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})

	require.NoError(t, err)

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	require.True(t, ok)

	assert.Equal(t, email, claims["email"].(string))
}
