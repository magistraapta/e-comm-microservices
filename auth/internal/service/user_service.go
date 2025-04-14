package service

import (
	"auth/internal/model"
	"auth/internal/repository"
	"auth/internal/utils"
	"context"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.Register(ctx, user)
}

func (s *UserService) Login(ctx context.Context, email string, password string) (*model.User, error) {
	existingUser, err := s.repo.UserLogin(ctx, email)

	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	if !utils.CheckHashPassword(password, existingUser.Email) {
		return nil, errors.New("Invalid email or password")
	}

	return existingUser, nil
}

func (s *UserService) AdminRegister(ctx context.Context, admin *model.Admin) (*model.Admin, error) {
	return s.repo.RegisterAdmin(ctx, admin)
}
