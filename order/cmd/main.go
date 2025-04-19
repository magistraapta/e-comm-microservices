package main

import (
	"fmt"
	"log"
	"net"
	"order/internal/client"
	"order/internal/handler"
	"order/internal/initializers"
	"order/internal/pb"

	"google.golang.org/grpc"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	db := initializers.ConnectDatabase()

	productClient := client.NewProductServiceClient("localhost:50052")
	handler := handler.NewOrderHandler(db, productClient)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, handler)

	fmt.Println("Order service is running on port :50053")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
