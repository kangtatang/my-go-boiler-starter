package handler

import (
	"project/internal/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,email"` // Validasi email untuk username
	Password string `json:"password" validate:"required,min=6"`
	// Username string `json:"username"`
	// Password string `json:"password"`
}

// LoginResponse mendeskripsikan respons sukses untuk login
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse mendeskripsikan struktur respons error
type ErrorResponse struct {
	Error string `json:"error"`
}

// Define a custom response struct for the profile response
type ProfileResponse struct {
	Message string `json:"message"`
}

// @Summary Login
// @Description Authenticate user and return JWT token
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login Request"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/login [post]
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

		return c.JSON(LoginResponse{Token: t})
	}
}

// @Summary Profile
// @Description Get user profile information
// @Produce json
// @Success 200 {object} ProfileResponse
// @Router /api/profile [get]
func Profile(c *fiber.Ctx) error {
	return c.JSON(ProfileResponse{Message: "Welcome to your profile!"})
}
