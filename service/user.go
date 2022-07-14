package service

import (
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type UserService interface {
	FindByEmail(email string) (model.User, error)
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

func (us *userService) FindByEmail(email string) (model.User, error) {
	user, err := us.userRepository.FindByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *userService) CreateUser(user model.User) (model.User, error) {
	user, err := us.userRepository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
