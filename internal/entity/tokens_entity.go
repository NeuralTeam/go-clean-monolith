package entity

import (
	"time"
)

// Tokens entity
type Tokens struct {
	Base
	User      Users     `json:"-"`
	UserID    int64     `gorm:"type:integer;index;" json:"-"`
	Token     string    `gorm:"type:varchar(64);unique;" json:"token"`
	CreatedAt time.Time `gorm:"type:timestamp;index;autoCreateTime;" json:"created_at"`
}

// TableName gives table name of entity
func (t *Tokens) TableName() string {
	return "tokens"
}
