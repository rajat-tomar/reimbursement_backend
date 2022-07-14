package service

import (
	"log"
	"reimbursement_backend/model"
)

type OAuthService interface {
	Login(user model.User) error
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

func (oauth *oauthService) Login(user model.User) error {
	_, err := oauth.FindByEmail(user.Email)
	if err != nil {
		_, err = oauth.CreateUser(user)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (oauth *oauthService) FindByEmail(email string) (model.User, error) {
	user, err := oauth.userService.FindByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (oauth *oauthService) CreateUser(user model.User) (model.User, error) {
	createdUser, err := oauth.userService.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}
