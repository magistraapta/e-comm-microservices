package main

import (
	"api-gateaway/internal/auth"
	"api-gateaway/internal/product"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	authConn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer authConn.Close()

	productConn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	authClient := auth.NewAuthServiceClient(authConn)
	authHandler := auth.NewAuthHandler(authClient)

	productClient := product.NewProductClient(productConn)
	productHandler := product.NewProductHandler(productClient)

	r := gin.Default()

	auth.SetupAuthRoute(r, authHandler)
	product.SetupProductRoute(r, productHandler)

	r.Run(":8000")
}
