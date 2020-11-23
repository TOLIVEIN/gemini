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
	Username  string `validate:"required,min=1,max=20"`
	Password  string `validate:"required,min=6"`
	Email     string `validate:"required,email"`
}

//Auth ...
type Auth struct {
	ID       uint   `gorm:"primarykey"`
	Username string `validate:"required,min=1,max=20"`
	Password string `validate:"required,min=6"`
}

//Article ...
type Article struct {
	// gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// TagID       uint   `gorm:"foreignKey"`
	Tag         Tag    `gorm:"embedded"`
	Title       string `validate:"required,max=100"`
	Description string `validate:"max=255"`
	Content     string
	CreatedBy   string `validate:"max=20"`
	UpdatedBy   string
}

//Tag ...
type Tag struct {
	// gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `validate:"required,max=20"`
	CreatedBy string `validate:"max=20"`
	UpdatedBy string
}
