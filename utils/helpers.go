package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateRandomState() string {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte("rajat")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
