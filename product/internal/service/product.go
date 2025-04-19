package service

import (
	"context"
	"errors"
	"product/internal/model"
	"product/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	if product.Name == "" || product.Price <= 0 || product.Stock <= 0 {
		return nil, errors.New("invalid product data")
	}

	return s.repo.CreateProduct(ctx, product)
}

func (s *ProductService) FindOne(ctx context.Context, product *model.Product) (*model.Product, error) {
	return s.repo.FindOne(ctx, product)
}

func (s *ProductService) FindAll(ctx context.Context) ([]model.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) DecreaseStock(ctx context.Context, productID int64, orderID int64, quantity int64) (*model.Product, error) {

	product, err := s.FindOne(ctx, &model.Product{Id: productID})

	if err != nil {
		return nil, err
	}

	if product.Stock <= 0 {
		return nil, errors.New("product is out of stock")
	}

	return s.repo.DecreaseStock(ctx, productID, orderID, quantity)
}
