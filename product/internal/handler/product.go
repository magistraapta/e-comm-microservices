package handler

import (
	"context"
	"net/http"
	"product/internal/model"
	"product/internal/pb"
	"product/internal/service"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product model.Product

	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock

	res, err := h.service.CreateProduct(ctx, &product)

	if err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusOK,
		Id:     res.Id,
	}, nil

}

func (h *ProductHandler) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product model.Product

	product.Id = req.Id
	res, err := h.service.FindOne(ctx, &product)

	if err != nil {
		return &pb.FindOneResponse{
			Status: http.StatusBadRequest,
			Error:  "failed to get product",
		}, nil
	}

	data := &pb.FindOneData{
		Id:    res.Id,
		Name:  res.Name,
		Price: res.Price,
		Stock: res.Stock,
	}

	return &pb.FindOneResponse{
		Data: data,
	}, nil
}

func (h *ProductHandler) FindAll(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {

	products, err := h.service.FindAll(ctx)

	if err != nil {
		return &pb.FindAllResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	var pbProducts []*pb.FindOneData

	for _, p := range products {
		pbProducts = append(pbProducts, &pb.FindOneData{
			Id:    p.Id,
			Name:  p.Name,
			Price: p.Price,
			Stock: p.Stock,
		})
	}

	return &pb.FindAllResponse{
		Products: pbProducts,
	}, nil
}
