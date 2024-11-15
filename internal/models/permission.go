package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
