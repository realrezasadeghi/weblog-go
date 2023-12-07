package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"weblog/constants"
	"weblog/models"
)

func GenerateToken(email, role string, id uint) (string, error) {
	fmt.Println("[GenerateToken] Generating access token")

	nowTime := time.Now()
	expiresTime := nowTime.Add(time.Minute * time.Duration(constants.JwtAccessTokenTimeDuration))

	tokenClaims := models.Token{
		Id:    id,
		Role:  role,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: constants.JwtAccessTokenIssuer,
			IssuedAt: &jwt.NumericDate{
				Time: nowTime,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: expiresTime,
			},
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString([]byte(constants.JwtAccessTokenSecret))

	if err != nil {
		fmt.Println("[GenerateTokenHelper]", err.Error())
		return "", err
	}

	return accessToken, nil
}

func ValidateToken(signedToken string) (*models.Token, error) {
	fmt.Println("[ValidateToken] Validating access token")

	token, err := jwt.ParseWithClaims(signedToken, &models.Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			message := fmt.Sprintln("invalid token ", token.Header["alg"])
			fmt.Println("[ValidateToken]", message)
			return nil, errors.New(message)
		}
		return []byte(constants.JwtAccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("[ValidateToken]", err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*models.Token)

	if !ok {
		errMessage := constants.ErrInvalidToken
		fmt.Println("[ValidateToken]", errMessage)
		return nil, errors.New(errMessage)
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		errMessage := constants.ErrExpiredToken
		fmt.Println("[ValidateToken]", errMessage)
		return nil, errors.New(errMessage)
	}

	fmt.Println("[ValidateTokenHelper] Validated access token")

	return claims, nil
}
