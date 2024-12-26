package userService

type UserService struct {
	repo userRepository
}

func NewUserService(r userRepository) *UserService {
	return &UserService{
		repo: r,
	}
}
