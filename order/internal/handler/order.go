package handler

import (
	"context"
	"net/http"
	"order/internal/client"
	"order/internal/model"
	"order/internal/pb"

	"gorm.io/gorm"
)

type OrderHandler struct {
	db *gorm.DB
	pb.UnimplementedOrderServiceServer
	ProductClient *client.ProductServiceClient
}

func NewOrderHandler(db *gorm.DB, productClient *client.ProductServiceClient) *OrderHandler {
	return &OrderHandler{
		db:            db,
		ProductClient: productClient,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := h.ProductClient.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	}

	if product.Data.Stock <= 0 {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  "stock is too low",
		}, nil
	}

	order := model.Order{
		Price:     product.Data.Price,
		ProductID: product.Data.Id,
		UserID:    req.UserId,
	}

	if err := h.db.Create(&order).Error; err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	res, err := h.ProductClient.DecreaseStock(order.ProductID, order.Id, req.Quantity)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	if res.Status == http.StatusConflict {
		h.db.Delete(&model.Order{}, order.Id)
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  res.Error,
		}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
