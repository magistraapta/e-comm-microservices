proto:
	protoc --go_out=. --go-grpc_out=. ./internal/pb/order.proto
	protoc --go_out=. --go-grpc_out=. ./internal/pb/product.proto


server:
	go run cmd/main.go
