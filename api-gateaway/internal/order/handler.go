package order

import (
	"api-gateaway/internal/order/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Client *OrderClient
}

func NewOrderHandler(client *OrderClient) *OrderHandler {
	return &OrderHandler{Client: client}
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req struct {
		ProductID int64 `json:"product_id"`
		Quantity  int64 `json:"quantity"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	val, exists := ctx.Get("userId")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userId, ok := val.(int64)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user id format"})
		return
	}

	res, err := h.Client.Client.CreateOrder(ctx, &pb.CreateOrderRequest{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		UserID:    userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
