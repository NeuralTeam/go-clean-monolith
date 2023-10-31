package entity

// Base model
type Base struct {
	ID int64 `gorm:"type:integer;primary_key;" json:"id"`
}
