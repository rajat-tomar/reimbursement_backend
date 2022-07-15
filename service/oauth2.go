package service

import (
	"fmt"
	"reimbursement_backend/model"
)

type OAuthService interface {
	Login(user model.User) (string, error)
	FindByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
}

type oauthService struct {
	userService UserService
}

func NewOAuthService() *oauthService {
	return &oauthService{
		userService: NewUserService(),
	}
}

func (oauth *oauthService) Login(user model.User) (string, error) {
	usr, err := oauth.FindByEmail(user.Email)
	if err != nil {
		usr, err = oauth.CreateUser(user)
		if err != nil {
			return "", fmt.Errorf("failed to login %v", err)
		}
	}

	return usr.Role, nil
}

func (oauth *oauthService) FindByEmail(email string) (model.User, error) {
	user, err := oauth.userService.FindByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to find user %v", err)
	}

	return user, nil
}

func (oauth *oauthService) CreateUser(user model.User) (model.User, error) {
	createdUser, err := oauth.userService.CreateUser(user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user %v", err)
	}

	return createdUser, nil
}
