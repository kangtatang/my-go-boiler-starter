package routes

import (
	"project/internal/handler"
	"project/internal/middleware"
	"project/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App, cfg *config.Config) {
	api := app.Group("/api")

	api.Post("/login", handler.Login(cfg.App.JWTSecret))
	api.Get("/profile", middleware.JWTProtected(cfg.App.JWTSecret), handler.Profile)
}
