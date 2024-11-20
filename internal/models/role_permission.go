package models

import (
	"time"
)

type RolePermission struct {
	RoleID       uint
	PermissionID uint
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
