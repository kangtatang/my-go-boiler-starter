package handler

import (
	"project/internal/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(jwtSecret string, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginReq LoginRequest
		if err := c.BodyParser(&loginReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}

		// Cek apakah user dengan username ini ada di database
		user, err := userService.GetUserByUsername(loginReq.Username)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}

		// Verifikasi password dengan bcrypt
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}

		// Buat token JWT
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["sub"] = user.ID
		claims["username"] = user.Username
		claims["role"] = user.Role
		claims["permissions"] = user.Permissions
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token kadaluarsa dalam 72 jam

		t, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not generate token"})
		}

		// Set role and permissions ke dalam context
		c.Locals("userRole", user.Role)
		c.Locals("userPermissions", user.Permissions)

		return c.JSON(fiber.Map{"token": t})
	}
}

func Profile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Welcome to your profile!"})
}
