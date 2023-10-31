package entity

import (
	"time"
)

// Users entity
type Users struct {
	Base
	Username  string    `gorm:"type:varchar(32);unique;not null;" json:"username"`
	Email     string    `gorm:"type:varchar(64);unique;not null;" json:"email"`
	Password  string    `gorm:"type:varchar(32);not null;" json:"password"`
	IsActive  bool      `gorm:"type:boolean;default:false;" json:"is_active"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime;" json:"created_at"`
}

// TableName gives table name of entity
func (t *Users) TableName() string {
	return "users"
}
