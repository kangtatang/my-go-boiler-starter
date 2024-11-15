package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID uuid.UUID `gorm:"type:uuid"`
	RoleID uint
}
