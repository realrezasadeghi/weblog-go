package utils

import (
	"net/http"
	"weblog/models"
	"weblog/responses"
)

func CreateSuccessResponse(code int, message string, data interface{}) responses.Success {
	return responses.Success{
		Code:    code,
		Data:    data,
		Message: message,
		Status:  http.StatusText(code),
	}
}

func CreateErrorResponse(code int, message string) responses.Error {
	return responses.Error{
		Code:    code,
		Message: message,
		Status:  http.StatusText(code),
	}
}

func CreateUserResponse(user models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:        user.ID,
		Role:      user.Role,
		Email:     user.Email,
		LastName:  user.LastName,
		FirstName: user.FirstName,
	}
}
