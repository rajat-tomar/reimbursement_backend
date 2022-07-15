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
}

func NewOAuthController() *oauthController {
	return &oauthController{
		oauthService: service.NewOAuthService(),
	}
}

func (oauth *oauthController) Login(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var requestUser model.User
	name := r.Context().Value("name")
	email := r.Context().Value("email")
	requestUser.Name = name.(string)
	requestUser.Email = email.(string)

	role, err := oauth.oauthService.Login(requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	response.Data = map[string]string{"role": role}
	_ = json.NewEncoder(w).Encode(response)
}
