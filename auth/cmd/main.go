package main

import (
	"auth/internal/handler"
	"auth/internal/initializer"
	"auth/internal/pb"
	"auth/internal/repository"
	"auth/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDatabase()
}

func main() {
	listener, err := net.Listen("tcp", ":5001")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	db := initializer.ConnectDatabase()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHanlder(userService)

	pb.RegisterAuthServiceServer(grpcServer, userHandler)

	log.Println("AuthService is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
