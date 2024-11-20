package handler

import (
	"net/http"
	"project/internal/models"
	"project/internal/service"

	myValidator "project/internal/utils/validator"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UserHandler - Struct untuk handler user
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler - Fungsi untuk membuat instance baru dari UserHandler
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// GetAllUsersResponse - Struct untuk response GetAllUsers
type GetAllUsersResponse struct {
	Data  []models.User `json:"data"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

type CreateUserResponse struct {
	Message string      `json:"message"`
	Data    models.User `json:"data"`
}

type UpdateUserResponse struct {
	Message string      `json:"message"`
	Data    models.User `json:"data"`
}

// type ErrorResponse struct {
// 	Error string `json:"error"`
// }

// @Summary Get all users
// @Description Retrieve a list of users with pagination, filtering, and sorting
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of users per page" default(10)
// @Param sort query string false "Sorting criteria" default("created_at desc")
// @Param username query string false "Filter by username"
// @Success 200 {object} GetAllUsersResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)              // Default page 1
	limit := c.QueryInt("limit", 10)           // Default limit 10
	sort := c.Query("sort", "created_at desc") // Default sorting by creation date desc
	filter := make(map[string]interface{})

	// Filter by username if provided
	username := c.Query("username")
	if username != "" {
		filter["username"] = username
	}

	// Call service to get users data
	users, total, err := h.userService.GetAllUsers(page, limit, sort, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get users"})
	}

	// Respond with user data and pagination info
	return c.JSON(GetAllUsersResponse{
		Data:  users, // Ensure users is of type []models.User
		Total: int(total),
		Page:  page,
		Limit: limit,
	})
}

// GetUserByID - Retrieve a user by their ID
// @Summary Get user by ID
// @Description Retrieve a user by their ID
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	// Call service to get user by ID
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// CreateUserRequest - Request body structure for creating a new user
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required"`
}

// @Summary Create a new user
// @Description Create a new user with the provided details
// @Accept json
// @Produce json
// @Param createUserRequest body CreateUserRequest true "Create User Request"
// @Success 201 {object} handler.CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate the request using a custom validator
	if err := myValidator.ValidateStruct(&req); err != nil {
		var errorMessage string
		var errorDetails = make(map[string]string)

		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldErr := range validationErrors {
				switch fieldErr.Field() {
				case "Username":
					errorDetails["username"] = "Username must be a valid email"
				case "Password":
					errorDetails["password"] = "Password must be at least 6 characters"
				case "Role":
					errorDetails["role"] = "Role is required"
				default:
					errorDetails[fieldErr.Field()] = "Invalid input"
				}
			}
			errorMessage = "Validation errors occurred"
		} else {
			errorMessage = "An unknown validation error occurred"
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorMessage,
			"errors":  errorDetails,
		})
	}

	// Hash the password
	hashedPassword, err := service.HashPassword(req.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create a new user
	newUser := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
	}

	// Call the service to create the user
	if err := h.userService.CreateUser(&newUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(http.StatusCreated).JSON(map[string]interface{}{
		"message": "User created successfully",
		"data":    newUser,
	})
}

// EditUserRequest - Request body structure for editing a user
type EditUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password,omitempty" validate:"min=6"`
	Role     string `json:"role" validate:"required"`
}

// @Summary Update an existing user
// @Description Update user details by their ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param editUserRequest body EditUserRequest true "Edit User Request"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var req EditUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate the request
	if err := myValidator.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// Fetch the existing user
	existingUser, err := h.userService.GetUserByID(userID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Update fields
	if req.Username != "" {
		existingUser.Username = req.Username
	}

	if req.Role != "" {
		existingUser.Role = req.Role
	}

	// Hash password if provided
	if req.Password != "" {
		hashedPassword, err := service.HashPassword(req.Password)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}
		existingUser.Password = hashedPassword
	}

	// Call service to update user
	if err := h.userService.UpdateUser(userID, &existingUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "User updated successfully",
		"data":    existingUser,
	})
}

// @Summary Delete a user
// @Description Delete a user by their ID
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} nil
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	// Call service to delete user
	if err := h.userService.DeleteUser(userID); err != nil {
		if err == service.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to delete user"})
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
