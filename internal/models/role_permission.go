package models

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleID       uint
	PermissionID uint
}
