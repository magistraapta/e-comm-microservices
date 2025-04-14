package handler

import (
	"auth/internal/model"
	"auth/internal/pb"
	"auth/internal/service"
	"auth/internal/utils"
	"context"
	"net/http"
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
		Status: http.StatusCreated,
		Id:     res.Id,
	}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.service.Login(ctx, user.Email, user.Password)

	if err != nil {
		return &pb.AuthResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	token, _ := utils.GenerateToken(res.Id)

	return &pb.AuthResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (h *AuthHandler) AdminRegister(ctx context.Context, req *pb.AdminRegisterResponse) (*pb.RegisterResponse, error) {
	admin := &model.Admin{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := h.service.AdminRegister(ctx, admin)

	if err != nil {
		return &pb.RegisterResponse{
			Error: err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Id:     res.Id,
	}, nil
}
