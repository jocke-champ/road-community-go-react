package models

import (
	"github.com/jinzhu/gorm"
)

//Add surname+lastname in User as Name
type User struct {
	gorm.Model
	Email        string     `gorm:"type:varchar(100);unique_index" json:"email"`
	PasswordHash string     `json:"password"`
	Name         string     `json:"name"`
	Posts        []Post     `gorm:"foreignkey:UserID"`
	Comments     []Comment  `gorm:"foreignkey:UserID"`
	Locations    []Location `gorm:"foreignkey:UserID"`
}

type Post struct {
	gorm.Model
	Content  string    `json:"content"`
	UserID   uint      `json:"userId"`
	Comments []Comment `gorm:"foreignkey:PostID"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
}

type Location struct {
	gorm.Model
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Description string  `json:"description"`
	UserID      uint    `json:"userId"`
}
