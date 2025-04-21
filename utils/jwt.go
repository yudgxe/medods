package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const key string = "oem2LsejuU3o0292ljsu"

var ErrInvalidToken = errors.New("invalid token")

func CreateToken(claims jwt.MapClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(key))
}

func VerificationToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	}

	return nil, ErrInvalidToken
}
