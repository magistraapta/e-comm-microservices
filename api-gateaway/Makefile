proto:
	protoc --go_out=. --go-grpc_out=. ./internal/auth/pb/auth.proto
	protoc --go_out=. --go-grpc_out=. ./internal/product/pb/product.proto
	protoc --go_out=. --go-grpc_out=. ./internal/order/pb/order.proto

server:
	go run cmd/main.go