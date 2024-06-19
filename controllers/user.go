package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"weblog/constants"
	"weblog/requests"
	"weblog/responses"
	"weblog/services"
	"weblog/utils"

	"github.com/gin-gonic/gin"
)

func handleErrorResponse(c *gin.Context, statusCode int, err error) {
	fmt.Printf("[%s] %s\n", c.HandlerName(), err.Error())
	errRes := utils.CreateErrorResponse(statusCode, err.Error())
	c.JSON(statusCode, errRes)
	return
}

type IUserController interface {
	LoginController(c *gin.Context)
	SignupController(c *gin.Context)
	GetUserController(c *gin.Context)
	UpdateUserController(c *gin.Context)
	DeleteUserController(c *gin.Context)
	GetAllUsersController(c *gin.Context)
	ForgetPasswordController(c *gin.Context)
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
		handleErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := uc.us.LoginUser(request)

	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, err)
	}

	token, errGenerateToken := utils.GenerateToken(user.Email, user.Role, user.ID)

	if errGenerateToken != nil {
		handleErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseUser := map[string]string{"token": token, "email": user.Email, "first_name": user.FirstName, "last_name": user.LastName, "role": user.Role}

	utils.SetAccessTokenCookie(c, token)
	c.JSON(http.StatusOK, utils.CreateSuccessResponse(http.StatusOK, constants.LoginSuccessful, responseUser))
}

func (uc *UserController) SignupController(c *gin.Context) {
	var request *requests.Signup

	if err := c.ShouldBindJSON(&request); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, err)
	}

	createdUser, err := uc.us.SignupUser(request)

	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseUser := map[string]string{"email": createdUser.Email, "first_name": createdUser.FirstName, "last_name": createdUser.LastName, "role": createdUser.Role}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.SignupSuccessful, responseUser)

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) GetUserController(c *gin.Context) {
	userId, isExistId := c.Get(constants.Id)

	if !isExistId {
		errMessage := constants.ErrInvalidId
		handleErrorResponse(c, http.StatusBadRequest, errors.New(errMessage))
	}

	user, err := uc.us.GetUserById(utils.ToString(userId))

	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.GetUserByEmailSuccessful, utils.CreateUserResponse(user))
	c.JSON(http.StatusOK, response)
}

func (uc *UserController) DeleteUserController(c *gin.Context) {
	userId, isExistId := c.Get(constants.Id)

	if !isExistId {
		errMessage := constants.ErrInvalidId
		handleErrorResponse(c, http.StatusBadRequest, errors.New(errMessage))
	}

	err := uc.us.DeleteUserById(utils.ToString(userId))

	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.DeleteUserByIdSuccessful, nil)

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) UpdateUserController(c *gin.Context) {
	id, isExistEmail := c.Get(constants.Id)

	if !isExistEmail {
		errMessage := constants.ErrInvalidToken
		handleErrorResponse(c, http.StatusInternalServerError, errors.New(errMessage))
	}

	var request *requests.Update

	if err := c.ShouldBindJSON(&request); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := uc.us.UpdateUserById(utils.ToString(id), request)

	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.UpdateSuccessful, utils.CreateUserResponse(user))

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) GetAllUsersController(c *gin.Context) {
	allUsers, err := uc.us.GetAllUsers()

	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, err)
	}
	var users []responses.UserResponse

	for _, user := range allUsers {
		users = append(users, utils.CreateUserResponse(user))
	}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.GetAllUsersSuccessful, users)
	c.JSON(http.StatusOK, response)
}

func (uc *UserController) ForgetPasswordController(c *gin.Context) {
	var request *requests.ForgetPassword

	if err := c.ShouldBindJSON(&request); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	errForgetPassword := uc.us.ForgetPasswordUser(request)

	if errForgetPassword != nil {
		handleErrorResponse(c, http.StatusInternalServerError, errForgetPassword)
		return
	}

	response := utils.CreateSuccessResponse(http.StatusOK, constants.ForgetPasswordSuccessful, nil)

	c.JSON(http.StatusOK, response)
}
