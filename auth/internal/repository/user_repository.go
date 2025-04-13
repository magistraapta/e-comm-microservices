package repository

import (
	"auth/internal/model"
	"auth/internal/utils"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Register(ctx context.Context, user *model.User) (*model.User, error) {

	// check email
	var existingUser model.User

	result := r.DB.WithContext(ctx).Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		return nil, errors.New("User already exist")
	}
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	newUser := model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	result = r.DB.WithContext(ctx).Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
