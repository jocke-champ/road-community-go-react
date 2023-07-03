package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email		string 'gorm:"type:varchar(100);unique_index" json:"email"'
	PasswordHash	string 'json:"password"'
}
