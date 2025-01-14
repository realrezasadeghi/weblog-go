package models

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	jwt.RegisteredClaims
	Id    uint   `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
