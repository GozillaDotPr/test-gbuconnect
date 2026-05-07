package main

import (
	"log"

	"github.com/app/gin-postgres-api/internal/config"
	"github.com/app/gin-postgres-api/internal/container"
	"github.com/app/gin-postgres-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	c := container.New(cfg)

	r := gin.Default()
	routes.Register(r, c.AuthHandler, c.ProductHandler, c.AuthService)

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
