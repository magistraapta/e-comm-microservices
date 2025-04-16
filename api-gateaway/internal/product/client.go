package product

import (
	"api-gateaway/internal/product/pb"

	"google.golang.org/grpc"
)

type ProductClient struct {
	Client pb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		Client: pb.NewProductServiceClient(conn),
	}
}
