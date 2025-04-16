package repository

import (
	"context"
	"errors"
	"product/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {

	newProduct := model.Product{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}

	err := r.db.WithContext(ctx).Create(&newProduct).Error

	if err != nil {
		return nil, err
	}

	return &newProduct, nil
}

func (r *ProductRepository) FindOne(ctx context.Context, product *model.Product) (*model.Product, error) {
	var existingProduct model.Product
	err := r.db.WithContext(ctx).First(&existingProduct, product.Id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		return nil, err
	}

	return &existingProduct, nil
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	// Use context-aware DB operation
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

