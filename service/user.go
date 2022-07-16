package service

import (
	"fmt"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type UserService interface {
	Login(user model.User) (string, error)
	FindByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUsers() ([]model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService() *userService {
	return &userService{
		userRepository: repository.NewUserRepository(),
	}
}

func (us *userService) Login(user model.User) (string, error) {
	loginUser, err := us.FindByEmail(user.Email)
	if err != nil {
		loginUser, err = us.CreateUser(user)
		if err != nil {
			return "", fmt.Errorf("failed to login %v", err)
		}
	}

	return loginUser.Role, nil
}

func (us *userService) FindByEmail(email string) (model.User, error) {
	user, err := us.userRepository.FindByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to find user %v", err)
	}

	return user, nil
}

func (us *userService) CreateUser(user model.User) (model.User, error) {
	createdUser, err := us.userRepository.CreateUser(user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user %v", err)
	}

	return createdUser, nil
}

func (us *userService) GetUsers() ([]model.User, error) {
	users, err := us.userRepository.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get users %v", err)
	}

	return users, nil
}
