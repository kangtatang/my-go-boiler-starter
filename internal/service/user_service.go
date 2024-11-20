package service

import (
	"errors"
	"project/internal/models"
	"project/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error)
	GetUserByID(id string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
	FindUserByID(id string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error) {
	// Memanggil repository untuk mendapatkan semua user dengan filter, pagination, dan sorting
	users, total, err := s.repo.GetAllUsers(page, limit, sort, filter)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (s *userService) GetUserByID(id string) (models.User, error) {
	// Memanggil repository untuk mendapatkan user berdasarkan ID
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername(username string) (models.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

// HashPassword melakukan hashing terhadap password user
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// FindUserByID retrieves a user by their ID
func (s *userService) FindUserByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id string, user *models.User) error {
	// Use the repository to update the user in the database
	if err := s.repo.UpdateUser(id, user); err != nil {
		return err
	}
	return nil
}

// Definisikan ErrUserNotFound
var ErrUserNotFound = errors.New("user not found")

func (s *userService) DeleteUser(id string) error {
	// Memastikan user ada sebelum menghapus
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.ID == uuid.Nil {
		// return errors.New("user not found")
		return ErrUserNotFound
	}

	// Memanggil repository untuk menghapus user
	if err := s.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
