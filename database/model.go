package database

import (
	"time"
)

//User ...
type User struct {
	// gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Password  string
	Email     string
}

//Article ...
type Article struct {
	// gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// TagID       uint   `gorm:"foreignKey"`
	Tag         Tag    `gorm:"embedded"`
	Title       string `validate:"required,alphanumunicode"`
	Description string `validate:"alphanumunicode"`
	Content     string
	CreatedBy   string `validate:"alphanumunicode"`
	UpdatedBy   string
}

//Tag ...
type Tag struct {
	// gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `validate:"required,alphanumunicode"`
	CreatedBy string `validate:"alphanumunicode"`
	UpdatedBy string
}
