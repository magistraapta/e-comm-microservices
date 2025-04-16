package product

import (
	"github.com/gin-gonic/gin"
)

func SetupProductRoute(r *gin.Engine, productHandler *ProductHandler) {

	product := r.Group("/product")
	// product.Use(authMiddleware.ValidateToken)
	product.POST("/", productHandler.CreateProduct)
	product.GET("/:id", productHandler.FindOne)
	product.GET("/", productHandler.FindAll)

}
