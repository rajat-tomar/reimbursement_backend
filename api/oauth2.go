package api

import (
	"encoding/json"
	"net/http"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
)

type OAuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type oauthController struct {
	oauthService service.OAuthService
	userService  service.UserService
}

func NewOAuthController() *oauthController {
	return &oauthController{
		oauthService: service.NewOAuthService(),
		userService:  service.NewUserService(),
	}
}

func (oauth *oauthController) Login(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var requestUser model.User
	name := r.Context().Value("name")
	email := r.Context().Value("email")
	requestUser.Name = name.(string)
	requestUser.Email = email.(string)

	err := oauth.oauthService.Login(requestUser)
	if err != nil {
		response.Message = "Failed to login"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response.Message = "Login successful"
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}
