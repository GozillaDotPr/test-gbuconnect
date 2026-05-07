package container

import (
	"github.com/app/gin-postgres-api/internal/config"
	"github.com/app/gin-postgres-api/internal/handler"
	"github.com/app/gin-postgres-api/internal/models"
	"github.com/app/gin-postgres-api/internal/repository"
	"github.com/app/gin-postgres-api/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	DB             *gorm.DB
	AuthService    service.AuthService
	ProductService service.ProductService
	AuthHandler    handler.AuthHandler
	ProductHandler handler.ProductHandler
}

func New(cfg *config.Config) *Container {
	db := config.NewDB(cfg)

	// Auto migrate
	db.AutoMigrate(&models.Product{})

	// Wire up
	productRepo := repository.NewProductRepository(db)
	authSvc := service.NewAuthService(cfg)
	productSvc := service.NewProductService(productRepo)

	return &Container{
		DB:             db,
		AuthService:    authSvc,
		ProductService: productSvc,
		AuthHandler:    handler.NewAuthHandler(authSvc),
		ProductHandler: handler.NewProductHandler(productSvc),
	}
}
