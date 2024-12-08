package main

import (
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello today we are going to build jwt authentication system.")
	err := godotenv.Load("../dev.env")
	if err != nil {
		slog.Error("Failed to load the env file", slog.String("error - ", err.Error()))
	}
	tokenString, err := CreateToken("rahul", "rahul@mail.com", "789")
	if err != nil {
		slog.Error("Error while creating the token", slog.String("error - ", err.Error()))
	}
	fmt.Println(tokenString)

	fmt.Println("trying to verify the token - ")
	err = VerifyToken(tokenString)
	fmt.Println(err)
}
