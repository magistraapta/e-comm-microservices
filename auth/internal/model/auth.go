package model

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Admin struct {
	Id       int64 `gorm:"primaryKey"`
	Username string
	Password string
}
