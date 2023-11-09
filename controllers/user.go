package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"weblog/constants"
	"weblog/requests"
	"weblog/responses"
	"weblog/services"
	"weblog/utils"
)

type IUserController interface {
	LoginController(c *gin.Context)
	SignupController(c *gin.Context)
	GetAllUsersController(c *gin.Context)
	GetUserByEmailController(c *gin.Context)
}

type UserController struct {
	us services.IUserService
}

func CreateUserController(us services.IUserService) IUserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) LoginController(c *gin.Context) {
	var request *requests.Login

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("[LoginController]", err.Error())
		errRes := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := uc.us.Login(request)

	if err != nil {
		fmt.Println("[LoginController]", err.Error())
		errRes := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	token, errGenerateToken := utils.GenerateToken(user.Email, user.Role)

	if errGenerateToken != nil {
		fmt.Println("[LoginController]", err.Error())
		errRes := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	responseUser := map[string]string{"token": token, "email": user.Email, "first_name": user.FirstName, "last_name": user.LastName, "role": user.Role}

	utils.SetAccessTokenCookie(c, token)
	c.JSON(http.StatusOK, utils.CreateSuccessResponse(http.StatusOK, constants.LoginSuccessful, responseUser))
}

func (uc *UserController) SignupController(c *gin.Context) {
	var request *requests.Signup

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("[SignupController]", err.Error())
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	createdUser, err := uc.us.Signup(request)

	if err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	responseUser := map[string]string{"email": createdUser.Email, "first_name": createdUser.FirstName, "last_name": createdUser.LastName, "role": createdUser.Role}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.SignupSuccessful, responseUser)

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) GetAllUsersController(c *gin.Context) {
	allUsers, err := uc.us.GetAllUsers()

	if err != nil {
		fmt.Println("[GetAllUsersController]", err.Error())
		errRes := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	var users []responses.UserResponse

	for _, user := range allUsers {
		users = append(users, utils.CreateUserResponse(user))
	}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.GetAllUsersSuccessful, users)
	c.JSON(http.StatusOK, response)
}

func (uc *UserController) GetUserByEmailController(c *gin.Context) {
	email := c.Param(constants.Email)

	if email == "" {
		errMessage := constants.ErrInvalidEmail
		fmt.Println("[GetUserByEmailController]", errMessage)
		errRes := utils.CreateErrorResponse(http.StatusBadRequest, errMessage)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := uc.us.GetUserByEmail(email)

	if err != nil {
		fmt.Println("[UserByUsernameHandler]", err.Error())
		errRes := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.GetUserByEmailSuccessful, user)
	c.JSON(http.StatusOK, response)
}
