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
	authorization, err := ctx.Cookie("Authorization")
	if err != nil || authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization cookie"})
		return
	}

	// No need to split or trim Bearer
	token := strings.TrimSpace(authorization)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is empty"})
		return
	}

	res, err := s.client.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token,
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.Set("userId", res.UserID)
	ctx.Next()
}
