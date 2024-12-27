package userService

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}
func (r *UserService) CreateUser(user User) (User, error) {
	hashedPassword, err := r.HashPassword(user.Password)
	log.Printf("service: %v", user)
	if err != nil {
		return User{}, err
	}
	user.Password = hashedPassword
	return r.repo.CreateUser(user)
}
func (r *UserService) DeleteUserById(id uint) (User, error) {
	return r.repo.DeleteUserById(id)
}
func (r *UserService) GetAllUsers() ([]User, error) {
	return r.repo.GetAllUsers()
}
func (r *UserService) UpdateUserById(id uint, user User) (User, error) {
	return r.repo.UpdateUserById(id, user)
}

func (r *UserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}
