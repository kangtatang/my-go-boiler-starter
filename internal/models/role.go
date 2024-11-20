package models

import "time"

type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
