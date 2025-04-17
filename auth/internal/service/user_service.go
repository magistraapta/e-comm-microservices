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
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckHashPassword(password, existingUser.Password) {
		return nil, errors.New("invalid email or password")
	}

	return existingUser, nil
}

func (s *UserService) AdminLogin(ctx context.Context, username string) (*model.Admin, error) {
	existingAdmin, err := s.repo.AdminLogin(ctx, username)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return existingAdmin, nil
}

func (s *UserService) AdminRegister(ctx context.Context, admin *model.Admin) (*model.Admin, error) {
	return s.repo.RegisterAdmin(ctx, admin)
}
