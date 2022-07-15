package api

import (
	"encoding/json"
	"net/http"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
)

type UserController interface {
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
