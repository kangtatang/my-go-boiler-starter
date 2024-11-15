package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWTProtected memvalidasi token JWT dan mengekstrak klaim user
func JWTProtected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mendapatkan token dari header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or malformed JWT"})
		}

		tokenString := authHeader[7:] // Menghilangkan "Bearer " dari header

		// Parsing dan validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		// Jika ada error atau token tidak valid
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired JWT"})
		}

		// Ekstrak klaim jika token valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Simpan klaim yang dibutuhkan ke dalam context untuk diakses di handler berikutnya
			c.Locals("userId", claims["userId"])
			c.Locals("userRole", claims["role"]) // misalkan klaim "role" digunakan untuk peran user
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
		}

		// Jika valid, lanjutkan ke handler berikutnya
		return c.Next()
	}
}
