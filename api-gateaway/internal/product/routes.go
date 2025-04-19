package product

import (
	"api-gateaway/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupProductRoute(r *gin.Engine, productHandler *ProductHandler, authSvc *auth.AuthMiddleware) {

	product := r.Group("/product")

	product.Use(authSvc.ValidateToken)
	product.GET("/:id", productHandler.FindOne)
	product.GET("/", productHandler.FindAll)

	product.POST("/", productHandler.CreateProduct)

}
