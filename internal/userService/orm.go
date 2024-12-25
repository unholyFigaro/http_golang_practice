package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
