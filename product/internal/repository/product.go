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

func (r *ProductRepository) DecreaseStock(ctx context.Context, productID int64, orderID int64, quantity int64) (*model.Product, error) {
	var product model.Product

	// 1. Find product by ID
	if err := r.db.WithContext(ctx).First(&product, productID).Error; err != nil {
		return nil, err
	}

	// 2. Check if stock decrease log already exists
	var existingLog model.StockDecreaseLog
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&existingLog).Error
	if err == nil {
		// Log already exists, don't decrease stock again
		return nil, errors.New("stock has already been decreased for this order")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Actual DB error
		return nil, err
	}

	// 3. Make sure stock is not zero or negative
	if product.Stock <= 0 {
		return nil, errors.New("not enough stock")
	}

	// 4. Decrease stock
	product.Stock -= quantity
	if err := r.db.WithContext(ctx).Save(&product).Error; err != nil {
		return nil, err
	}

	// 5. Create stock decrease log
	log := model.StockDecreaseLog{
		OrderID:      orderID,
		ProductRefer: productID,
	}
	if err := r.db.WithContext(ctx).Create(&log).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
