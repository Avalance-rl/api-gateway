package gateway

import (
	authv1 "github.com/avalance-rl/otiva/proto/gen/avalance.auth.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	authClient authv1.AuthServiceClient
}

func NewGateway(authAddr string) (*Gateway, error) {
	conn, err := grpc.NewClient(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Gateway{
		authClient: authv1.NewAuthServiceClient(conn),
	}, nil
}
