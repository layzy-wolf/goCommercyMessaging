package suite

import (
	authv1 "app/api/protobuf/auth.v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strconv"
	"testing"
	"time"
)

type Suite struct {
	*testing.T
	AuthClient authv1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Duration.Hours(1)))

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	p, _ := strconv.Atoi("440444")

	grpcAddress := net.JoinHostPort("localhost", strconv.Itoa(p))

	cc, err := grpc.DialContext(context.Background(),
		grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	authClient := authv1.NewAuthClient(cc)

	return ctx, &Suite{
		T:          t,
		AuthClient: authClient,
	}
}
