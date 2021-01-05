package database

import (
	"time"

	"gorm.io/gorm"
)

//User ...
type User struct {
	// gorm.Model `json:"id,createdAt,updatedAt,deletedAt"`
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `validate:"required,min=1,max=20" json:"username"`
	Password  string         `validate:"required,min=6" json:"password"`
	Email     string         `validate:"required,email" json:"email"`
}

//Auth ...
type Auth struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `validate:"required,min=1,max=20" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}

//Article ...
type Article struct {
	// gorm.Model `json:"id,createdAt,updatedAt,deletedAt"`
	// TagIDs        []uint
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Title         string         `validate:"required,max=100" json:"title"`
	Description   string         `validate:"max=255" json:"description"`
	CoverImageURL string         `validate:"required,max=255" json:"coverImageURL"`
	Content       string         `json:"content"`
	CreatedBy     string         `validate:"max=20" json:"createdBy"`
	UpdatedBy     string         `json:"updatedBy"`
	Tags          []*Tag         `gorm:"many2many:article_tag" json:"tags"`
}

//Tag ...
type Tag struct {
	// gorm.Model `json:"id,createdAt,updatedAt,deletedAt"`
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `validate:"required,max=20" json:"name"`
	CreatedBy string         `validate:"max=20" json:"createdBy"`
	UpdatedBy string         `json:"updatedBy"`
	Articles  []*Article     `gorm:"many2many:article_tag" json:"articles"`
}
