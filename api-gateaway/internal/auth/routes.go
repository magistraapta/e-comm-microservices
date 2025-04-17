package auth

import "github.com/gin-gonic/gin"

func SetupAuthRoute(r *gin.Engine, authHandler *AuthHandler) {

	auth := r.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	admin := r.Group("/admin")
	admin.POST("/register", authHandler.AdminRegister)
	admin.POST("/login", authHandler.AdminLogin)

}
