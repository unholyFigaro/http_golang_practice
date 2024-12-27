package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserById(id uint, user User) (User, error)
	DeleteUserById(id uint) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users)
	if err.Error != nil {
		return []User{}, err.Error
	}
	return users, nil
}

func (r *userRepository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user)
	if err.Error != nil {
		return User{}, err.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUserById(id uint) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return User{}, nil
	}
	if err := r.db.Model(&user).Delete(&user, id); err != nil {
		return User{}, nil
	}
	return user, nil
}

func (r *userRepository) UpdateUserById(id uint, user User) (User, error) {
	var existingUser User
	err := r.db.Model(&user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return User{}, err
	}
	r.db.Find(&existingUser, id)
	return existingUser, nil
}
