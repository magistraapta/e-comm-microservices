package auth

import (
	"api-gateaway/internal/auth/pb"

	"google.golang.org/grpc"
)

type AuthClient struct {
	Client pb.AuthServiceClient
}

func NewAuthServiceClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		Client: pb.NewAuthServiceClient(conn),
	}
}
