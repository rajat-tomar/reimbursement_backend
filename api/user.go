package api

import (
	"encoding/json"
	"net/http"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func NewUserController() *userController {
	return &userController{
		userService: service.NewUserService(),
	}
}

func (uc *userController) Login(w http.ResponseWriter, r *http.Request) {
	var requestUser model.User
	name := r.Context().Value("name")
	email := r.Context().Value("email")
	requestUser.Name = name.(string)
	requestUser.Email = email.(string)

	role, err := uc.userService.Login(requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	_ = json.NewEncoder(w).Encode(role)
}

func (uc *userController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	_ = json.NewEncoder(w).Encode(users)
}
