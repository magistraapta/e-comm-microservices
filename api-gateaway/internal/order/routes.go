package order

import (
	"api-gateaway/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoute(r *gin.Engine, orderHandler *OrderHandler, authSvc *auth.AuthMiddleware) {
	order := r.Group("/order")
	order.Use(authSvc.ValidateToken)
	order.POST("/", orderHandler.CreateOrder)
}
