package api

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"reimbursement_backend/config"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type OAuthController interface {
	GoogleLogin(w http.ResponseWriter, r *http.Request)
	GoogleCallback(w http.ResponseWriter, r *http.Request)
}

type oauthController struct {
	googleOauthConfig *oauth2.Config
	oauthStateString  string
}

func NewOAuthController() *oauthController {
	return &oauthController{
		googleOauthConfig: &oauth2.Config{
			ClientID:     config.Configuration.OAuth.Google.ClientID,
			ClientSecret: config.Configuration.OAuth.Google.ClientSecret,
			Scopes:       config.Configuration.OAuth.Google.Scopes,
			RedirectURL:  config.Configuration.OAuth.Google.RedirectURL,
			Endpoint:     google.Endpoint,
		},
	}
}

func (oauth *oauthController) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauth.oauthStateString = oauth.generateOAuthState(w)
	url := oauth.googleOauthConfig.AuthCodeURL(oauth.oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (oauth *oauthController) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != oauth.oauthStateString {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := oauth.getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func (oauth *oauthController) generateOAuthState(w http.ResponseWriter) string {
	//var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	//cookie := http.Cookie{Name: "oauthState", Value: state, Expires: expiration, Secure: true, HttpOnly: true}
	//http.SetCookie(w, &cookie)

	return state
}

func (oauth *oauthController) getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := oauth.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	return contents, nil
}
