package routes

import (
	"project/internal/handler"
	"project/internal/middleware"
	"project/internal/repository"
	"project/internal/service"
	"project/pkg/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// InitializeRoutes mengatur semua route dan middleware yang dibutuhkan
func InitializeRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	// Group untuk API utama
	api := app.Group("/api")

	// Inisialisasi komponen User (repository, service, handler)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Route untuk autentikasi dan profil, mengirimkan userService ke Login
	api.Post("/login", handler.Login(cfg.App.JWTSecret, userService))
	api.Get("/profile", middleware.JWTProtected(cfg.App.JWTSecret), handler.Profile)

	// Group untuk route user yang membutuhkan autentikasi dan otorisasi admin
	userRoutes := api.Group("/users", middleware.JWTProtected(cfg.App.JWTSecret))

	// Routes untuk User dengan akses role admin
	userRoutes.Get("/", middleware.RequireRole("superadmin"), userHandler.GetAllUsers)
	userRoutes.Get("/:id", middleware.RequireRole("superadmin"), userHandler.GetUserByID)
	userRoutes.Post("/", middleware.RequireRole("superadmin"), userHandler.CreateUser)
	userRoutes.Put("/:id", middleware.RequireRole("superadmin"), userHandler.UpdateUser)
	userRoutes.Delete("/:id", middleware.RequireRole("superadmin"), userHandler.DeleteUser)

	// Jika Anda ingin menggunakan permission-based access di masa depan:
	// userRoutes.Get("/:id", middleware.RequirePermission("view_user", userService), userHandler.GetUserByID)
	// userRoutes.Post("/", middleware.RequirePermission("create_user", userService), userHandler.CreateUser)
	// userRoutes.Put("/:id", middleware.RequirePermission("edit_user", userService), userHandler.UpdateUser)
	// userRoutes.Delete("/:id", middleware.RequirePermission("delete_user", userService), userHandler.DeleteUser)
}

// package routes

// import (
// 	"project/internal/handler"
// 	"project/internal/middleware"
// 	"project/internal/repository"
// 	"project/internal/service"
// 	"project/pkg/config"

// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/gorm"
// )

// // InitializeRoutes mengatur semua route dan middleware yang dibutuhkan
// func InitializeRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
// 	// Group untuk API utama
// 	api := app.Group("/api")

// 	// Inisialisasi komponen User (repository, service, handler)
// 	userRepository := repository.NewUserRepository(db)
// 	userService := service.NewUserService(userRepository)
// 	userHandler := handler.NewUserHandler(userService)

// 	// Route untuk autentikasi dan profil, mengirimkan userService ke Login
// 	api.Post("/login", handler.Login(cfg.App.JWTSecret, userService))
// 	api.Get("/profile", middleware.JWTProtected(cfg.App.JWTSecret), handler.Profile)

// 	// Group untuk route user yang membutuhkan autentikasi dan otorisasi admin
// 	userRoutes := api.Group("/users", middleware.JWTProtected(cfg.App.JWTSecret))
// 	// userRoutes.Get("/", middleware.RequireAdminRole(), userHandler.GetAllUsers)
// 	// userRoutes.Get("/:id", middleware.RequireAdminRole(), userHandler.GetUserByID)
// 	// userRoutes.Post("/", middleware.RequireAdminRole(), userHandler.CreateUser)
// 	// userRoutes.Put("/:id", middleware.RequireAdminRole(), userHandler.UpdateUser)
// 	// userRoutes.Delete("/:id", middleware.RequireAdminRole(), userHandler.DeleteUser)

// 	// User routes, memerlukan role admin
// 	//  userRoutes := api.Group("/users", middleware.JWTProtected(cfg.App.JWTSecret))
// 	userRoutes.Get("/", middleware.RequireRole("admin"), userHandler.GetAllUsers)
// 	userRoutes.Get("/:id", middleware.RequireRole("admin"), userHandler.GetUserByID)
// 	userRoutes.Post("/", middleware.RequireRole("admin"), userHandler.CreateUser)
// 	userRoutes.Put("/:id", middleware.RequireRole("admin"), userHandler.UpdateUser)
// 	userRoutes.Delete("/:id", middleware.RequireRole("admin"), userHandler.DeleteUser)
// }
