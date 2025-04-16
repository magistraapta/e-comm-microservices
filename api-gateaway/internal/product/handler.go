package product

import (
	"api-gateaway/internal/product/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Client *ProductClient
}

func NewProductHandler(client *ProductClient) *ProductHandler {
	return &ProductHandler{Client: client}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Price int64  `json:"price"`
		Stock int64  `json:"stock"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.Client.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	if res == nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "nil response from product service"})
		return
	}

	ctx.JSON(http.StatusOK, &pb.CreateProductResponse{
		Id:     res.Id,
		Status: http.StatusOK,
	})
}

func (h *ProductHandler) FindOne(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to get id"})
		return
	}

	res, err := h.Client.Client.FindOne(ctx, &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "failed to get product"})
		return
	}

	ctx.JSON(http.StatusOK, &pb.FindOneResponse{
		Data: res.Data,
	})
}

func (h *ProductHandler) FindAll(ctx *gin.Context) {
	products, err := h.Client.Client.FindAll(ctx, &pb.FindAllRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}
