package models

import (
	"time"
)

type Permission struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
