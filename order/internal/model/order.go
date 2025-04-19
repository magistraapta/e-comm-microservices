package model

type Order struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	ProductID int64 `json:"product_id"`
	Price     int64 `json:"price"`
	UserID    int64 `json:"user_id"`
}
