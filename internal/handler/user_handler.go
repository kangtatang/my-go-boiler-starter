package handler

import (
	"net/http"
	"project/internal/models"
	"project/internal/service"
	"project/internal/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	// Ambil parameter paginasi, filter, dan sorting dari query params
	page := c.QueryInt("page", 1)              // Default page 1
	limit := c.QueryInt("limit", 10)           // Default limit 10
	sort := c.Query("sort", "created_at desc") // Default sorting by creation date desc
	filter := make(map[string]interface{})

	// Misalnya, filter berdasarkan username jika ada
	username := c.Query("username")
	if username != "" {
		filter["username"] = username
	}

	// Panggil service untuk mendapatkan data user
	users, total, err := h.userService.GetAllUsers(page, limit, sort, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get users"})
	}

	// Response dengan data users dan pagination info
	return c.JSON(fiber.Map{
		"data":  users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	// Ambil ID dari parameter dan validasi
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	// Panggil service GetUserByID
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Mendefinisikan struktur request
	type CreateUserRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,min=6"`
		Role     string `json:"role" validate:"required"`
	}

	// Parsing request JSON ke dalam struct
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// // Validasi input
	// if err := c.App().Validator().Struct(&req); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Validation failed",
	// 	})
	// }

	// Validate the request struct
	if err := validator.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// Hashing password
	hashedPassword, err := service.HashPassword(req.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Membuat objek user baru
	newUser := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
	}

	// Memanggil service untuk membuat user di database
	if err := h.userService.CreateUser(&newUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    newUser,
	})

	// // Parse request body
	// var user models.User
	// if err := c.BodyParser(&user); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	// }

	// // Panggil service untuk membuat user baru
	// if err := h.userService.CreateUser(&user); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	// }

	// return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	// Ambil ID dari parameter dan validasi
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	// Parse request body
	var updatedData models.User
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	// Panggil service untuk mengupdate user
	if err := h.userService.UpdateUser(id, &updatedData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	// Ambil ID dari parameter dan validasi
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	// Panggil service untuk menghapus user
	if err := h.userService.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

// package handler

// import (
// 	"project/internal/service"

// 	"github.com/gofiber/fiber/v2"
// )

// type UserHandler struct {
// 	userService service.UserService
// }

// func NewUserHandler(userService service.UserService) *UserHandler {
// 	return &UserHandler{userService}
// }

// func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
// 	// Ambil parameter paginasi, filter, dan sorting dari query params
// 	// Panggil service untuk mendapatkan data user
// }

// func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
// 	// Ambil ID dari parameter dan panggil service GetUserByID
// }

// func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
// 	// Parse request body, validasi, lalu panggil service CreateUser
// }

// func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
// 	// Ambil ID dari parameter, parse request body, lalu panggil service UpdateUser
// }

// func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
// 	// Ambil ID dari parameter lalu panggil service DeleteUser
// }
