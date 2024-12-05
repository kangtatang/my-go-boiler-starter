package main

import (
	"log"
	_ "project/docs"
	"project/internal/routes"
	"project/internal/utils/validator"
	"project/pkg/config"
	"project/pkg/database"
	"project/pkg/seeder"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Sat Net Base User Management API
// @version 1.0
// @description This is a User Management API.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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

	// userRepo := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

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

	// Menambahkan middleware CORS {tambahkan url FE jika sudah ada nanti misalkan http://localhost:3000}
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080,http://127.0.0.1:8080", // Ganti dengan URL yang Anda gunakan untuk frontend
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin,Content-Type,Authorization",
	}))

	// // Inisialisasi validator
	// app.Validator = validator.NewValidator()
	// Initialize the validator
	validator.InitValidator()

	// Tambahkan middleware logger untuk mencatat semua request
	app.Use(logger.New())

	// Initialize routes dan injeksikan dependensi
	routes.InitializeRoutes(app, db, cfg)

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Start server
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
