package main

import (
	"log"
	"project/internal/routes"
	"project/internal/utils/validator"
	"project/pkg/config"
	"project/pkg/database"
	"project/pkg/seeder"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Inisialisasi koneksi database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Migrate the database
	if err := database.MigrateDatabase(db); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Seeder untuk membuat user super admin
	if err := seeder.SeedSuperAdmin(db, "passwordadmin"); err != nil {
		log.Fatalf("Error seeding super admin: %v", err)
	}

	// Inisialisasi Fiber
	app := fiber.New()

	// // Inisialisasi validator
	// app.Validator = validator.NewValidator()
	// Initialize the validator
	validator.InitValidator()

	// Tambahkan middleware logger untuk mencatat semua request
	app.Use(logger.New())

	// Initialize routes dan injeksikan dependensi
	routes.InitializeRoutes(app, db, cfg)

	// Start server
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
