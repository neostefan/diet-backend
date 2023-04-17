package util

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userId int) (string, error) {

	var secret string = "The cow says moo!!"

	//! Add an expiration time to the jwt token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})

	jwtToken, err := t.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
