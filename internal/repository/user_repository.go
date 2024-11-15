package repository

import (
	"project/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error)
	GetUserByID(id string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// GetAllUsers memanggil database untuk mendapatkan semua pengguna
func (r *userRepository) GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// query := r.db.Model(&models.User{})
	// Membuat query dengan filter dan sort
	query := r.db.Model(&models.User{}).Preload("Roles").Preload("Permissions")

	// Apply filtering
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	// Apply sorting
	if sort != "" {
		query = query.Order(sort)
	}

	// Apply pagination
	query = query.Offset((page - 1) * limit).Limit(limit)

	// Count total number of records
	query.Count(&total)

	// Get users
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// GetUserByID memanggil database untuk mendapatkan pengguna berdasarkan ID
func (r *userRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return user, err
	}

	if err := r.db.First(&user, "id = ?", parsedID).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

// CreateUser menambahkan pengguna baru ke database
func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// UpdateUser memperbarui informasi pengguna di database
func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

// DeleteUser menghapus pengguna berdasarkan ID
func (r *userRepository) DeleteUser(id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&models.User{}, "id = ?", parsedID).Error
}

// package repository

// import (
// 	"fmt"
// 	"project/internal/models"

// 	"gorm.io/gorm"
// )

// // UserRepository adalah interface untuk repositori user
// type UserRepository interface {
// 	GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error)
// 	GetUserByID(id string) (models.User, error)
// 	CreateUser(user *models.User) error
// 	UpdateUser(user *models.User) error
// 	DeleteUser(id string) error
// }

// // userRepository adalah implementasi dari UserRepository
// type userRepository struct {
// 	db *gorm.DB
// }

// // NewUserRepository adalah constructor untuk membuat instansi baru userRepository
// func NewUserRepository(db *gorm.DB) UserRepository {
// 	return &userRepository{db}
// }

// // GetAllUsers untuk mengambil semua user dengan pagination, sorting, dan filtering
// func (r *userRepository) GetAllUsers(page, limit int, sort string, filter map[string]interface{}) ([]models.User, int64, error) {
// 	var users []models.User
// 	var total int64

// 	// Menghitung jumlah total data
// 	if err := r.db.Model(&models.User{}).Count(&total).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// Membuat query dasar untuk mengambil data
// 	query := r.db.Model(&models.User{})

// 	// Mengaplikasikan filter jika ada
// 	for key, value := range filter {
// 		query = query.Where(fmt.Sprintf("%s = ?", key), value)
// 	}

// 	// Menambahkan sorting
// 	if sort != "" {
// 		query = query.Order(sort)
// 	} else {
// 		// Jika tidak ada sorting, menggunakan urutan default berdasarkan ID
// 		query = query.Order("id desc")
// 	}

// 	// Mengaplikasikan pagination
// 	offset := (page - 1) * limit
// 	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	return users, total, nil
// }

// // GetUserByID untuk mengambil user berdasarkan ID
// func (r *userRepository) GetUserByID(id string) (models.User, error) {
// 	var user models.User
// 	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// // CreateUser untuk menambahkan user baru ke database
// func (r *userRepository) CreateUser(user *models.User) error {
// 	if err := r.db.Create(user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // UpdateUser untuk memperbarui data user berdasarkan ID
// func (r *userRepository) UpdateUser(user *models.User) error {
// 	if err := r.db.Save(user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // DeleteUser untuk menghapus user berdasarkan ID
// func (r *userRepository) DeleteUser(id string) error {
// 	if err := r.db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
