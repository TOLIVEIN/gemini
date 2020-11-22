package database

import (
	"gorm.io/gorm"
)

//User ...
type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

//Article ...
type Article struct {
	gorm.Model
	TagID       uint   `validate:"required,number"`
	Tag         Tag    `validate:"required"`
	Title       string `validate:"required,alphanumunicode"`
	Description string `validate:"alphanumunicode"`
	Content     string
	CreatedBy   string `validate:"alphanumunicode"`
	UpdatedBy   string
}

//Tag ...
type Tag struct {
	gorm.Model
	Name      string `validate:"required,alphanumunicode"`
	CreatedBy string `validate:"alphanumunicode"`
	UpdatedBy string
}
