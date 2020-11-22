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
	Tag         Tag
	Title       string
	Description string
	Content     string
	CreatedBy   string
	UpdatedBy   string
	DeletedBy   string
}

//Tag ...
type Tag struct {
	gorm.Model
	Name      string
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
