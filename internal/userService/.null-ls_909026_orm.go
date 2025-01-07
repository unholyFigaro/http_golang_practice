package userService

import (
	"github.com/unxly/golang-pa/internal/web/tasks"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint         `json:"id"`
	Email    string       `gorm:"unique;not null" json:"email"`
	Password string       `json:"password"`
	Tasks    []tasks.Task `json:"tasks" gorm:"foreignKey:UserId"`
}
