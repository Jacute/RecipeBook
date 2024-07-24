package utils

import (
	"RecipeBookApi/config"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) string {
	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"username": username,
		},
	)
	tokenString, err := token.SignedString(config.JWT_PRIVATE_KEY)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWT_PUBLIC_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Can't verify token")
	}
	return token, nil
}
