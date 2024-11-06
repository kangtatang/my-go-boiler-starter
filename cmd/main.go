package main

import (
	"log"
	"project/internal/routes"
	"project/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Inisialisasi Fiber
	app := fiber.New()

	// Tambahkan middleware logger untuk mencatat semua request
	app.Use(logger.New())

	// Initialize routes
	routes.InitializeRoutes(app, cfg)

	// Start server
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
