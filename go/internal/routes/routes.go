package routes

import (
	"github.com/app/gin-postgres-api/internal/handler"
	"github.com/app/gin-postgres-api/internal/middleware"
	"github.com/app/gin-postgres-api/internal/service"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, authHandler handler.AuthHandler, productHandler handler.ProductHandler, authService service.AuthService) {
	r.POST("/api/v1/login", authHandler.Login)

	products := r.Group("/api/v1/products", middleware.JWTAuth(authService))
	{
		products.GET("", productHandler.GetAll)
		products.GET("/:id", productHandler.GetByID)
		products.POST("", productHandler.Create)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)
	}
}
