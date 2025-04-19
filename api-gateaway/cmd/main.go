package main

import (
	"api-gateaway/internal/auth"
	"api-gateaway/internal/order"
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

	orderConn, err := grpc.Dial(":50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	authClient := auth.NewAuthServiceClient(authConn)
	authHandler := auth.NewAuthHandler(authClient)

	productClient := product.NewProductClient(productConn)
	productHandler := product.NewProductHandler(productClient)

	orderClient := order.NewOrderServiceClient(orderConn)
	orderHandler := order.NewOrderHandler(orderClient)

	r := gin.Default()

	middleware := auth.NewAuthMiddleware(authClient)

	auth.SetupAuthRoute(r, authHandler)
	product.SetupProductRoute(r, productHandler, &middleware)
	order.SetupOrderRoute(r, orderHandler, &middleware)

	r.Run(":8000")
}
