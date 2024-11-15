package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`
	Username    string       `gorm:"unique" json:"username"`
	Password    string       `gorm:"unique;not null" json:"-"`
	Role        string       `gorm:"not null" json:"role"`
	Roles       []Role       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
