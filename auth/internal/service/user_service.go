package service

import (
	"auth/internal/model"
	"auth/internal/repository"
	"context"
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
