package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `json:"id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"password"`
}
