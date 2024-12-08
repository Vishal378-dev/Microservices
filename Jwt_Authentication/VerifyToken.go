package main

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("SECRET_KEY")
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("Error while parsing token")
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("unable to extract claims from token")
	}
	fmt.Println("claims after verifying the token - ", claims)
	return nil
}
