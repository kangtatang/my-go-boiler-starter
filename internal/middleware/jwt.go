package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

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

		// Jika valid, lanjutkan ke handler berikutnya
		return c.Next()
	}
}

// func JWTProtected(secret string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		tokenString := c.Get("Authorization")
// 		if tokenString == "" {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or malformed JWT"})
// 		}

// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(secret), nil
// 		})

// 		if err != nil || !token.Valid {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired JWT"})
// 		}

// 		return c.Next()
// 	}
// }

// func LogToElk(message string) {
// 	client, _ := elasticsearch.NewDefaultClient()
// 	// Log handling and error reporting here
// 	log.Println(message) // Placeholder for logging to ELK
// }
