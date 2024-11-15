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

// func (s *userService) CreateUser(user *models.User) error {
// 	// Memanggil repository untuk membuat user baru
// 	if err := s.repo.CreateUser(user); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

// HashPassword melakukan hashing terhadap password user
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *userService) UpdateUser(id string, user *models.User) error {
	// Memastikan user ada sebelum memperbarui
	existingUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	if existingUser.ID == uuid.Nil {
		return errors.New("user not found")
	}

	// Memanggil repository untuk memperbarui user
	user.ID = existingUser.ID // Pastikan ID tetap sama
	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteUser(id string) error {
	// Memastikan user ada sebelum menghapus
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.ID == uuid.Nil {
		return errors.New("user not found")
	}

	// Memanggil repository untuk menghapus user
	if err := s.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}

// package service

// import (
// 	"errors"
// 	"project/internal/models"
// 	"project/internal/repository"

// 	"github.com/google/uuid"
// )

// // UserService adalah interface yang mendefinisikan operasi terkait user
// type UserService interface {
// 	GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error)
// 	GetUserByID(id string) (models.User, error)
// 	CreateUser(user *models.User) error
// 	UpdateUser(id string, user *models.User) error
// 	DeleteUser(id string) error
// }

// // userService adalah struct yang mengimplementasikan UserService
// type userService struct {
// 	repo repository.UserRepository
// }

// // NewUserService adalah constructor untuk membuat instance baru dari userService
// func NewUserService(repo repository.UserRepository) UserService {
// 	return &userService{repo}
// }

// // GetAllUsers memanggil repository untuk mendapatkan data users dengan pagination, sorting, dan filter
// func (s *userService) GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error) {
// 	// Memanggil repository untuk mendapatkan semua user
// 	users, total, err := s.repo.GetAllUsers(page, limit, sort, filter)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	return users, total, nil
// }

// // GetUserByID memanggil repository untuk mendapatkan user berdasarkan ID
// func (s *userService) GetUserByID(id string) (models.User, error) {
// 	// Memanggil repository untuk mencari user berdasarkan ID
// 	user, err := s.repo.GetUserByID(id)
// 	if err != nil {
// 		return models.User{}, err
// 	}
// 	return user, nil
// }

// // CreateUser memanggil repository untuk membuat user baru
// func (s *userService) CreateUser(user *models.User) error {
// 	// Memanggil repository untuk menambahkan user baru
// 	if err := s.repo.CreateUser(user); err != nil {
// 		return err
// 	}
// 	return nil
// }

// // UpdateUser memanggil repository untuk memperbarui user berdasarkan ID
// func (s *userService) UpdateUser(id string, user *models.User) error {
// 	// Mengonversi id string menjadi uuid.UUID jika menggunakan UUID
// 	parsedID, err := uuid.Parse(id)
// 	if err != nil {
// 		return errors.New("invalid user ID format")
// 	}

// 	// Memastikan user ada sebelum melakukan update
// 	existingUser, err := s.repo.GetUserByID(parsedID.String()) // Ubah ke penggunaan ID UUID
// 	if err != nil {
// 		return err
// 	}
// 	if existingUser.ID == uuid.Nil {
// 		return errors.New("user not found")
// 	}

// 	// Memanggil repository untuk memperbarui user
// 	user.ID = existingUser.ID // Pastikan ID tetap sama
// 	if err := s.repo.UpdateUser(user); err != nil {
// 		return err
// 	}
// 	return nil
// }

// // DeleteUser memanggil repository untuk menghapus user berdasarkan ID
// func (s *userService) DeleteUser(id string) error {
// 	// Mengonversi id string menjadi uuid.UUID
// 	parsedID, err := uuid.Parse(id)
// 	if err != nil {
// 		return errors.New("invalid user ID format")
// 	}

// 	// Memastikan user ada sebelum dihapus
// 	user, err := s.repo.GetUserByID(parsedID.String()) // Menggunakan UUID sebagai string
// 	if err != nil {
// 		return err
// 	}
// 	if user.ID == uuid.Nil {
// 		return errors.New("user not found")
// 	}

// 	// Memanggil repository untuk menghapus user
// 	if err := s.repo.DeleteUser(parsedID.String()); err != nil {
// 		return err
// 	}
// 	return nil
// }
