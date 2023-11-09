package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	fmt.Println("[GenerateHashedPasswordHelper] Generating hash of the password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("[GenerateHashedPasswordHelper]", err.Error())
		return "", err
	}

	fmt.Println("[GenerateHashedPasswordHelper] Generated hash of the password")
	return string(hash), nil
}

func ComparePassword(hashedPass string, password string) bool {
	fmt.Println("[CheckPasswordHelper] Checking hashed password and input password")

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password)); err != nil {
		fmt.Println("[CheckPasswordHelper]", err.Error())
		return false
	}

	fmt.Println("[CheckPasswordHelper] Password is correct")
	return true
}
