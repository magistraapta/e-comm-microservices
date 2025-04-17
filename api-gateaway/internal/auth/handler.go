package auth

import (
	"api-gateaway/internal/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Client *AuthClient
}

func NewAuthHandler(client *AuthClient) *AuthHandler {
	return &AuthHandler{Client: client}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.Client.Register(context.Background(), &pb.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(int(res.Status), res)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.Client.Login(context.Background(), &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *AuthHandler) AdminRegister(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.Client.AdminRegister(ctx, &pb.AdminRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), res)
}

func (h *AuthHandler) AdminLogin(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.Client.AdminLogin(ctx, &pb.AdminLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), res)

}
