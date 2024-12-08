package main

import (
	"fmt"
	"regexp"
)

type CreateTokenBody struct {
	Username string
	Email    string
	Id       string
}

func (c *CreateTokenBody) ValidateRequestBody() (bool, error) {
	if c.Username == "" || c.Username == "null" || c.Username == "undefined" {
		return false, fmt.Errorf("invalid username")
	}
	if c.Id == "" || c.Id == "null" || c.Id == "undefined" {
		return false, fmt.Errorf("invalid id")
	}

	if c.Email == "" || !ValidateEmail(c.Email) {
		return false, fmt.Errorf("invalid email")
	}
	return true, nil

}

func ValidateEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	result := regexp.MustCompile(emailPattern)
	return result.MatchString(email)
}
