package seeder

import (
	"log"
	"project/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedSuperAdmin akan menambahkan user super admin ke database
func SeedSuperAdmin(db *gorm.DB, password string) error {
	// Cek apakah sudah ada user super admin
	var user models.User
	if err := db.Where("username = ?", "superadmin").First(&user).Error; err == nil {
		// Jika ditemukan, berarti superadmin sudah ada
		log.Println("Super Admin sudah ada!")
		return nil
	}

	// Seed Role Superadmin dan Admin
	superAdminRole := models.Role{Name: "superadmin"}
	if err := db.Create(&superAdminRole).Error; err != nil {
		return err
	}

	adminRole := models.Role{Name: "admin"}
	if err := db.Create(&adminRole).Error; err != nil {
		return err
	}

	// Tambahkan permissions ke database
	permissions := []models.Permission{
		{Name: "create_user"},
		{Name: "edit_user"},
		{Name: "delete_user"},
		{Name: "view_user"},
	}
	if err := db.Create(&permissions).Error; err != nil {
		return err
	}

	// Kaitkan semua permissions dengan superAdminRole
	if err := db.Model(&superAdminRole).Association("Permissions").Append(&permissions); err != nil {
		return err
	}

	// Membuat user superadmin dengan role dan permission yang sudah ada
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	superAdmin := models.User{
		Username:    "superadmin",
		Password:    string(hashedPassword),
		Role:        "superadmin",
		Roles:       []models.Role{superAdminRole}, // Set role superadmin
		Permissions: permissions,                   // Set permissions langsung
	}

	// Menyimpan user super admin ke dalam database
	if err := db.Create(&superAdmin).Error; err != nil {
		return err
	}

	// **Hapus** penambahan role dan permission berulang pada user superAdmin

	log.Println("Super Admin berhasil dibuat!")
	return nil
}

// package seeder

// import (
// 	"log"
// 	"project/internal/models"

// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// // SeedSuperAdmin akan menambahkan user super admin ke database
// func SeedSuperAdmin(db *gorm.DB, password string) error {
// 	// Cek apakah sudah ada user super admin
// 	var user models.User
// 	if err := db.Where("username = ?", "superadmin").First(&user).Error; err == nil {
// 		// Jika ditemukan, berarti superadmin sudah ada
// 		log.Println("Super Admin sudah ada!")
// 		return nil
// 	}

// 	// Seed Role Superadmin dan Admin
// 	superAdminRole := models.Role{Name: "superadmin"}
// 	if err := db.Create(&superAdminRole).Error; err != nil {
// 		return err
// 	}

// 	adminRole := models.Role{Name: "admin"}
// 	if err := db.Create(&adminRole).Error; err != nil {
// 		return err
// 	}

// 	// Tambahkan permissions ke database
// 	permissions := []models.Permission{
// 		{Name: "create_user"},
// 		{Name: "edit_user"},
// 		{Name: "delete_user"},
// 		{Name: "view_user"},
// 	}
// 	if err := db.Create(&permissions).Error; err != nil {
// 		return err
// 	}

// 	// Kaitkan semua permissions dengan superAdminRole
// 	if err := db.Model(&superAdminRole).Association("Permissions").Append(&permissions); err != nil {
// 		return err
// 	}

// 	// Membuat user superadmin dengan role dan permission yang sudah ada
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	superAdmin := models.User{
// 		Username:    "superadmin",
// 		Password:    string(hashedPassword),
// 		Role:        "superadmin",
// 		Roles:       []models.Role{superAdminRole},
// 		Permissions: permissions,
// 	}

// 	// Menyimpan user super admin ke dalam database
// 	if err := db.Create(&superAdmin).Error; err != nil {
// 		return err
// 	}

// 	// Mengaitkan role dan permission ke user superadmin
// 	if err := db.Model(&superAdmin).Association("Roles").Append(&superAdminRole); err != nil {
// 		return err
// 	}
// 	if err := db.Model(&superAdmin).Association("Permissions").Append(&permissions); err != nil {
// 		return err
// 	}

// 	log.Println("Super Admin berhasil dibuat!")
// 	return nil
// }

// // package seeder

// // import (
// // 	"log"
// // 	"project/internal/models"

// // 	"golang.org/x/crypto/bcrypt"
// // 	"gorm.io/gorm"
// // )

// // // SeedSuperAdmin akan menambahkan user super admin ke database
// // func SeedSuperAdmin(db *gorm.DB, password string) error {
// // 	// Cek apakah sudah ada user super admin
// // 	var user models.User
// // 	// if err := db.Where("role = ?", "superadmin").First(&user).Error; err == nil {
// // 	// 	log.Println("Super Admin sudah ada!")
// // 	// 	return nil
// // 	// }

// // 	if err := db.Where("username = ?", "superadmin").First(&user).Error; err == nil {
// // 		// Jika ditemukan, berarti superadmin sudah ada
// // 		log.Println("Super Admin sudah ada!")
// // 		return nil
// // 	}

// // 	// Seed Roles dan Permissions
// // 	// Role Superadmin
// // 	superAdminRole := models.Role{Name: "superadmin"}
// // 	if err := db.Create(&superAdminRole).Error; err != nil {
// // 		return err
// // 	}

// // 	// Role Admin
// // 	SuperAdminRole := models.Role{Name: "superadmin"}
// // 	if err := db.Create(&SuperAdminRole).Error; err != nil {
// // 		return err
// // 	}

// // 	adminRole := models.Role{Name: "admin"}
// // 	db.Create(&adminRole)

// // 	// Tambahkan permissions ke database
// // 	permissions := []models.Permission{
// // 		{Name: "create_user"},
// // 		{Name: "edit_user"},
// // 		{Name: "delete_user"},
// // 		{Name: "view_user"},
// // 	}
// // 	for _, permission := range permissions {
// // 		db.Create(&permission)
// // 	}

// // 	// Permissions
// // 	createUserPermission := models.Permission{Name: "create_user"}
// // 	editUserPermission := models.Permission{Name: "edit_user"}
// // 	deleteUserPermission := models.Permission{Name: "delete_user"}
// // 	viewUserPermission := models.Permission{Name: "view_user"}
// // 	if err := db.Create(&createUserPermission).Error; err != nil {
// // 		return err
// // 	}
// // 	if err := db.Create(&editUserPermission).Error; err != nil {
// // 		return err
// // 	}
// // 	if err := db.Create(&deleteUserPermission).Error; err != nil {
// // 		return err
// // 	}
// // 	if err := db.Create(&viewUserPermission).Error; err != nil {
// // 		return err
// // 	}

// // 	// Mengaitkan Permissions dengan Super Admin Role
// // 	if err := db.Model(&superAdminRole).Association("Permissions").Append(&createUserPermission, &editUserPermission, &deleteUserPermission, &viewUserPermission); err != nil {
// // 		return err
// // 	}

// // 	// Membuat user baru dengan role superadmin
// // 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	superAdmin := models.User{
// // 		Username:    "superadmin", // Username untuk superadmin
// // 		Password:    string(hashedPassword),
// // 		Role:        "superadmin", // Set role
// // 		Roles:       []models.Role{superAdminRole},
// // 		Permissions: permissions,
// // 	}

// // 	// Menyimpan user super admin ke dalam database
// // 	if err := db.Create(&superAdmin).Error; err != nil {
// // 		return err
// // 	}

// // 	// Menyematkan superadmin role ke user yang baru dibuat
// // 	if err := db.Model(&superAdmin).Association("Roles").Append(&superAdminRole); err != nil {
// // 		return err
// // 	}

// // 	// Menyematkan permission ke user
// // 	if err := db.Model(&superAdmin).Association("Permissions").Append(&createUserPermission, &editUserPermission, &deleteUserPermission, &viewUserPermission); err != nil {
// // 		return err
// // 	}

// // 	log.Println("Super Admin berhasil dibuat!")
// // 	return nil
// // }

// // // package seeder

// // // import (
// // // 	"log"
// // // 	"project/internal/models"

// // // 	"golang.org/x/crypto/bcrypt"
// // // 	"gorm.io/gorm"
// // // )

// // // // SeedSuperAdmin akan menambahkan user super admin ke database
// // // func SeedSuperAdmin(db *gorm.DB, password string) error {
// // // 	// Cek apakah sudah ada user super admin
// // // 	var user models.User
// // // 	if err := db.Where("role = ?", "superadmin").First(&user).Error; err == nil {
// // // 		log.Println("Super Admin sudah ada!")
// // // 		return nil
// // // 	}

// // // 	// Seed Roles dan Permissions
// // // 	// Role Superadmin
// // // 	superAdminRole := models.Role{Name: "superadmin"}
// // // 	db.Create(&superAdminRole)

// // // 	// Role Admin
// // // 	adminRole := models.Role{Name: "admin"}
// // // 	db.Create(&adminRole)

// // // 	// Permissions
// // // 	createUserPermission := models.Permission{Name: "create_user"}
// // // 	editUserPermission := models.Permission{Name: "edit_user"}
// // // 	deleteUserPermission := models.Permission{Name: "delete_user"}
// // // 	viewUserPermission := models.Permission{Name: "view_user"}
// // // 	db.Create(&createUserPermission)
// // // 	db.Create(&editUserPermission)
// // // 	db.Create(&deleteUserPermission)
// // // 	db.Create(&viewUserPermission)

// // // 	// Mengaitkan Permissions dengan Super Admin Role
// // // 	db.Model(&superAdminRole).Association("Permissions").Append(&createUserPermission, &editUserPermission, &deleteUserPermission, &viewUserPermission)

// // // 	// Membuat user baru dengan role superadmin
// // // 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// // // 	if err != nil {
// // // 		return err
// // // 	}

// // // 	superAdmin := models.User{
// // // 		Username: "superadmin", // Username untuk superadmin
// // // 		Password: string(hashedPassword),
// // // 		Role:     "superadmin", // Set role
// // // 		Permissions: []string{"create_user", "edit_user", "delete_user", "view_user"},
// // // 	}

// // // 	// Menyimpan user super admin ke dalam database
// // // 	if err := db.Create(&superAdmin).Error; err != nil {
// // // 		return err
// // // 	}

// // // 	// Menyematkan superadmin role ke user yang baru dibuat
// // // 	db.Model(&superAdmin).Association("Roles").Append(&superAdminRole)

// // // 	log.Println("Super Admin berhasil dibuat!")
// // // 	return nil
// // // }
