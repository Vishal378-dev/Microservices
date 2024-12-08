package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string, email string, id string) (string, error) {
	RequestBody := CreateTokenBody{
		Username: username,
		Email:    email,
		Id:       id,
	}
	_, err := RequestBody.ValidateRequestBody()
	if err != nil {
		fmt.Printf("Invalid Request Body:- %s\n", err.Error())
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"email":    email,
			"id":       id,
		})
	secret_key := os.Getenv("SECRET_KEY")
	if secret_key == "" {
		log.Fatal("invalid secret")
	}
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
