package main

import (
	"fmt"
	"log"
	"net"
	"product/internal/handler"
	"product/internal/initializer"
	"product/internal/pb"
	"product/internal/repository"
	"product/internal/service"

	"google.golang.org/grpc"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDatabase()
}
func main() {
	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	db := initializer.ConnectDatabase()
	repo := repository.NewProductRepository(db)
	service := service.NewProductService(repo)
	handler := handler.NewProductHandler(service)

	pb.RegisterProductServiceServer(grpcServer, handler)

	fmt.Print("Product service is running on port :50052...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: &v", err)
	}
}
