package seeder

import (
	"log"
	"project/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedSuperAdmin akan menambahkan user super admin ke database
func SeedSuperAdmin(db *gorm.DB, password string) error {
	// Cek apakah role "superadmin" sudah ada atau buat jika belum ada
	var superAdminRole models.Role
	if err := db.Where("name = ?", "superadmin").First(&superAdminRole).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			superAdminRole = models.Role{Name: "superadmin"}
			if err := db.Create(&superAdminRole).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Daftar permissions yang dibutuhkan
	permissionsList := []string{"create_user", "edit_user", "delete_user", "view_user"}
	var permissions []models.Permission

	// Loop untuk memeriksa setiap permission, insert jika tidak ada
	for _, permName := range permissionsList {
		var permission models.Permission
		if err := db.Where("name = ?", permName).First(&permission).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Buat permission baru jika belum ada
				permission = models.Permission{Name: permName}
				if err := db.Create(&permission).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
		permissions = append(permissions, permission)
	}

	// Kaitkan semua permissions dengan superAdminRole jika belum terkait
	if err := db.Model(&superAdminRole).Association("Permissions").Replace(&permissions); err != nil {
		return err
	}

	// Cek apakah user "superadmin" sudah ada atau buat jika belum ada
	var superAdminUser models.User
	if err := db.Where("username = ?", "superadmin@mail.com").First(&superAdminUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Hash password superadmin
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			// Buat user superadmin baru
			superAdminUser = models.User{
				Username: "superadmin@mail.com",
				Password: string(hashedPassword),
				Role:     "superadmin",
			}

			if err := db.Create(&superAdminUser).Error; err != nil {
				return err
			}

			// Kaitkan permissions dengan user superadmin
			if err := db.Model(&superAdminUser).Association("Permissions").Replace(&permissions); err != nil {
				return err
			}

			// Kaitkan role superadmin dengan user superadmin
			if err := db.Model(&superAdminUser).Association("Roles").Replace(&superAdminRole); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	log.Println("Super Admin dan data terkait berhasil dibuat atau sudah ada.")
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
// 	// Cek apakah role superadmin sudah ada
// 	var superAdminRole models.Role
// 	if err := db.Where("name = ?", "superadmin").First(&superAdminRole).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			superAdminRole = models.Role{Name: "superadmin"}
// 			if err := db.Create(&superAdminRole).Error; err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	}

// 	// Cek apakah role admin sudah ada
// 	var adminRole models.Role
// 	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			adminRole = models.Role{Name: "admin"}
// 			if err := db.Create(&adminRole).Error; err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	}

// 	// Tambahkan permissions jika belum ada
// 	permissions := []models.Permission{
// 		{Name: "create_user"},
// 		{Name: "edit_user"},
// 		{Name: "delete_user"},
// 		{Name: "view_user"},
// 	}

// 	// Pastikan permissions ada di database
// 	for _, permission := range permissions {
// 		var perm models.Permission
// 		if err := db.Where("name = ?", permission.Name).First(&perm).Error; err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				if err := db.Create(&permission).Error; err != nil {
// 					return err
// 				}
// 			} else {
// 				return err
// 			}
// 		}
// 	}

// 	// Kaitkan semua permissions dengan superAdminRole jika belum terkait
// 	if err := db.Model(&superAdminRole).Association("Permissions").Clear(); err != nil {
// 		return err
// 	}
// 	if err := db.Model(&superAdminRole).Association("Permissions").Append(&permissions); err != nil {
// 		return err
// 	}

// 	// Cek apakah user superadmin sudah ada
// 	var user models.User
// 	if err := db.Where("username = ?", "superadmin@mail.com").First(&user).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			// Hash password
// 			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 			if err != nil {
// 				return err
// 			}

// 			// Buat user superadmin baru
// 			superAdmin := models.User{
// 				Username:    "superadmin@mail.com",
// 				Password:    string(hashedPassword),
// 				Role:        "superadmin",
// 				Roles:       []models.Role{superAdminRole},
// 				Permissions: permissions,
// 			}

// 			if err := db.Create(&superAdmin).Error; err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	}

// 	log.Println("Super Admin dan data terkait berhasil dibuat atau sudah ada.")
// 	return nil
// }

// package seeder

// import (
// 	"log"
// 	"project/internal/models"

// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// // SeedSuperAdmin akan menambahkan user super admin ke database
// func SeedSuperAdmin(db *gorm.DB, password string) error {
// 	// Cek apakah sudah ada role superadmin
// 	var superAdminRole models.Role
// 	if err := db.Where("name = ?", "superadmin").First(&superAdminRole).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			superAdminRole = models.Role{Name: "superadmin"}
// 			if err := db.Create(&superAdminRole).Error; err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	}

// 	// Cek apakah sudah ada role admin
// 	var adminRole models.Role
// 	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			adminRole = models.Role{Name: "admin"}
// 			if err := db.Create(&adminRole).Error; err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	}

// 	// Tambahkan permissions jika belum ada
// 	permissions := []models.Permission{
// 		{Name: "create_user"},
// 		{Name: "edit_user"},
// 		{Name: "delete_user"},
// 		{Name: "view_user"},
// 	}

// 	for _, permission := range permissions {
// 		var perm models.Permission
// 		if err := db.Where("name = ?", permission.Name).First(&perm).Error; err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				if err := db.Create(&permission).Error; err != nil {
// 					return err
// 				}
// 			} else {
// 				return err
// 			}
// 		}
// 	}

// 	// Kaitkan semua permissions dengan superAdminRole jika belum terkait
// 	if err := db.Model(&superAdminRole).Association("Permissions").Append(&permissions); err != nil {
// 		return err
// 	}

// 	// Cek apakah user superadmin sudah ada
// 	var user models.User
// 	if err := db.Where("username = ?", "superadmin@mail.com").First(&user).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			// Hash password
// 			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 			if err != nil {
// 				return err
// 			}

// 			// Buat user superadmin baru
// 			superAdmin := models.User{
// 				Username:    "superadmin@mail.com",
// 				Password:    string(hashedPassword),
// 				Role:        "superadmin",
// 				Roles:       []models.Role{superAdminRole},
// 				Permissions: permissions,
// 			}

// 			if err := db.Create(&superAdmin).Error; err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	}

// 	log.Println("Super Admin dan data terkait berhasil dibuat atau sudah ada.")
// 	return nil
// }

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
// 	if err := db.Where("username = ?", "superadmin@mail.com").First(&user).Error; err == nil {
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
// 		Username:    "superadmin@mail.com",
// 		Password:    string(hashedPassword),
// 		Role:        "superadmin",
// 		Roles:       []models.Role{superAdminRole}, // Set role superadmin
// 		Permissions: permissions,                   // Set permissions langsung
// 	}

// 	// Menyimpan user super admin ke dalam database
// 	if err := db.Create(&superAdmin).Error; err != nil {
// 		return err
// 	}

// 	// **Hapus** penambahan role dan permission berulang pada user superAdmin

// 	log.Println("Super Admin berhasil dibuat!")
// 	return nil
// }
