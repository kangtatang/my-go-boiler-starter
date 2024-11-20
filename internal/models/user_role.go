package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	UserID    uuid.UUID `gorm:"type:uuid"`
	RoleID    uint
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
