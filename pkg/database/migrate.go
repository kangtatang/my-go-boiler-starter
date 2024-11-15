package database

import (
	"log"
	"project/internal/models"

	"gorm.io/gorm"
)

// MigrateDatabase creates tables for all models
func MigrateDatabase(db *gorm.DB) error {
	log.Println("Migrating database...")

	// Auto-migrate the User model
	if err := db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.UserRole{}, &models.RolePermission{}); err != nil {
		return err
	}

	log.Println("Database migration completed.")
	return nil
}
