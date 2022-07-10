package service

import (
	"fmt"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type UserService interface {
	CreateUser(user model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService() *userService {
	return &userService{
		userRepository: repository.NewUserRepository(),
	}
}

func (us *userService) CreateUser(user model.User) (model.User, error) {
	user, err := us.userRepository.FindByEmail(user.Email)
	if err == nil {
		return model.User{}, fmt.Errorf("user already exists %v", err)
	}

	user, err = us.userRepository.CreateUser(user)

	return user, err
}
