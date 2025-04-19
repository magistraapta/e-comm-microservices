package order

import (
	"api-gateaway/internal/order/pb"

	"google.golang.org/grpc"
)

type OrderClient struct {
	Client pb.OrderServiceClient
}

func NewOrderServiceClient(conn *grpc.ClientConn) *OrderClient {
	return &OrderClient{
		Client: pb.NewOrderServiceClient(conn),
	}
}
