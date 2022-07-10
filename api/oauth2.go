package api

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
	"reimbursement_backend/utils"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type OAuthController interface {
	GoogleLogin(w http.ResponseWriter, r *http.Request)
	GoogleCallback(w http.ResponseWriter, r *http.Request)
}

type oauthController struct {
	oauthService service.OAuthService
	userService  service.UserService
}

var (
	oauthStateString  string
	googleOauthConfig *oauth2.Config
)

func NewOAuthController() *oauthController {
	return &oauthController{
		oauthService: service.NewOAuthService(),
		userService:  service.NewUserService(),
	}
}

func (oauth *oauthController) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleOauthConfig := initializeGoogleConfig()
	url := googleOauthConfig.AuthCodeURL(utils.GenerateRandomState())
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (oauth *oauthController) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var response model.Response

	if r.FormValue("state") != oauthStateString {
		log.Println("invalid oauth google state")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err)
		response.Message = "Failed to get user data from google"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	_ = json.Unmarshal(data, &user)
	token, err := utils.GenerateJWT(user.Email)
	user, err = oauth.userService.CreateUser(user)
	if err != nil {
		response.Message = "Logged in successfully"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
		return
	}

	log.Println("user created")
	json.NewEncoder(w).Encode(token)
}

func initializeGoogleConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.Configuration.OAuth.Google.ClientID,
		ClientSecret: config.Configuration.OAuth.Google.ClientSecret,
		Scopes:       config.Configuration.OAuth.Google.Scopes,
		RedirectURL:  config.Configuration.OAuth.Google.RedirectURL,
		Endpoint:     google.Endpoint,
	}
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	return contents, nil
}
