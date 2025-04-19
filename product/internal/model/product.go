package model

type Product struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Stock int64  `json:"stock"`
}

type StockDecreaseLog struct {
	Id           int64 `json:"id" gorm:"primaryKey"`
	OrderID      int64 `json:"order_id"`
	ProductRefer int64 `json:"product_refer"`
}
