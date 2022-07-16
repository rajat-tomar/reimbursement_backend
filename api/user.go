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
	var response model.Response
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

	response.Data = map[string]string{"role": role}
	_ = json.NewEncoder(w).Encode(response)
}

func (uc *userController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	users, err := uc.userService.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	response.Data = users
	_ = json.NewEncoder(w).Encode(response)
}
