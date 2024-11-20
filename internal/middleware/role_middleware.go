package middleware

import (
	"project/internal/service"

	"github.com/gofiber/fiber/v2"
)

// RequireRole adalah middleware untuk memastikan user memiliki role tertentu
func RequireRole(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mengambil role user dari context
		userRole := c.Locals("userRole")
		if userRole != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}
		return c.Next()
	}
}

// RequirePermission adalah middleware untuk memeriksa apakah user memiliki permission tertentu
func RequirePermission(requiredPermission string, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil user ID dari context (misalnya, setelah JWT divalidasi)
		userID := c.Locals("userID").(string)

		// Ambil user dari database
		user, err := userService.GetUserByID(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
		}

		// Memeriksa apakah user memiliki permission yang diperlukan
		hasPermission := false
		for _, permission := range user.Permissions {
			if permission.Name == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.Next()
	}
}
