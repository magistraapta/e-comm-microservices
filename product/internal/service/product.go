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
