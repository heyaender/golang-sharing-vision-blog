package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET string

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func VerifyJwtToken(strToken string) (*Claims, error) {
	if JWT_SECRET = os.Getenv("JWT_SECRET"); JWT_SECRET == "" {
		log.Fatal("[ ERROR ] JWT_SECRET is not defined!\n")
	}

	key := []byte(JWT_SECRET)

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, fmt.Errorf("Invalid token signature")
		}
	}

	if !token.Valid {
		return claims, fmt.Errorf("Invalid token")
	}

	return claims, nil
}