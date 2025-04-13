package handler

import (
	"auth/internal/model"
	"auth/internal/pb"
	"auth/internal/service"
	"context"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *service.UserService
}

func NewUserHanlder(service *service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.service.Register(ctx, user)

	if err != nil {
		return &pb.RegisterResponse{
			Error: err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Id: res.Id,
	}, nil
}
