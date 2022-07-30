package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"reimbursement_backend/config"
	"time"
)

type JWTClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateJWT(email string, name string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 365)
	claims := &JWTClaim{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(config.Config.JWTKey))
	return
}

func validateToken(signedToken string) (name, email string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWTKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", "", err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", "", err
	}
	return claims.Name, claims.Email, nil
}
