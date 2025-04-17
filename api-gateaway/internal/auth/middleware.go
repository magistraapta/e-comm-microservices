package auth

import (
	"api-gateaway/internal/auth/pb"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	client *AuthClient
}

func NewAuthMiddleware(client *AuthClient) AuthMiddleware {
	return AuthMiddleware{client: client}
}

func (s *AuthMiddleware) ValidateToken(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := s.client.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userID", res.UserID)

	ctx.Next()
}
