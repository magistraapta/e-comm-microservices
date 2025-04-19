package client

import (
	"context"
	"log"
	"order/internal/pb"

	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func NewProductServiceClient(url string) *ProductServiceClient {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}

	return &ProductServiceClient{
		Client: pb.NewProductServiceClient(conn),
	}
}

func (c *ProductServiceClient) FindOne(productID int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productID,
	}
	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(productID int64, orderID int64, quantity int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:       productID,
		OrderId:  orderID,
		Quantity: quantity,
	}
	return c.Client.DecreaseStock(context.Background(), req)
}
