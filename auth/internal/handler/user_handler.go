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
	var user model.User

	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password

	res, err := h.service.Register(ctx, &user)

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

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	res, err := h.service.Login(ctx, req.Email, req.Password)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	matchPassword := utils.CheckHashPassword(req.Password, res.Password)

	if !matchPassword {
		return &pb.LoginResponse{
			Status: http.StatusBadGateway,
			Error:  "Invalid email or password",
		}, nil
	}

	token, _ := utils.GenerateToken(res)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (h *AuthHandler) AdminRegister(ctx context.Context, req *pb.AdminRegisterRequest) (*pb.RegisterResponse, error) {
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

func (h *AuthHandler) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.LoginResponse, error) {
	res, err := h.service.AdminLogin(ctx, req.Username)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	isPasswordMatch := utils.CheckHashPassword(req.Password, res.Password)

	if !isPasswordMatch {
		return &pb.LoginResponse{
			Status: http.StatusBadGateway,
			Error:  "Invalid email or password",
		}, nil
	}

	token, err := utils.GenerateAdminToken(res)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (h *AuthHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := utils.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserID: claims.Id,
	}, nil
}
