package userService

type UserRepository interface {
	CreateUser(user User) (User, error)
}
