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
	token := r.Context().Value("token")
	requestUser.Name = name.(string)
	requestUser.Email = email.(string)
	tokenString := token.(string)

	role, err := uc.userService.Login(requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	data := map[string]string{
		"token": tokenString,
		"role":  role,
	}

	_ = json.NewEncoder(w).Encode(data)
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
