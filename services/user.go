package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"weblog/constants"
	"weblog/models"
	"weblog/repositories"
	"weblog/requests"
	"weblog/utils"
)

type IUserService interface {
	DeleteUserById(id string) error
	GetAllUsers() ([]models.User, error)
	GetUserByEmail(email string) (models.User, error)
	Login(request *requests.Login) (*models.User, error)
	Signup(request *requests.Signup) (*models.User, error)
	Update(email string, request *requests.Update) (models.User, error)
}

type UserService struct {
	ur repositories.IUserRepository
}

func CreateUserService(ur repositories.IUserRepository) IUserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	fmt.Println("[GetAllUsersController] list of users if user service")

	users, err := us.ur.GetAllUsers()

	if err != nil {
		fmt.Println("[GetAllUsersController]", err.Error())
		return []models.User{}, err
	}

	fmt.Println("[GetAllUsersController] list of users has found in user service")

	return users, nil
}

func (us *UserService) DeleteUserById(id string) error {
	fmt.Println("[DeleteUserById] Hitting delete user by id in user service")

	err := us.ur.DeleteUserById(id)

	if err != nil {
		fmt.Println("[DeleteUserById]", err.Error())
		return err
	}

	fmt.Println("[DeleteUserById] user deleted from table in user service")

	return nil
}

func (us *UserService) GetUserByEmail(email string) (models.User, error) {
	fmt.Println("[GetUserByEmailPath] get details user by email in user service")

	user, err := us.ur.GetUserByEmail(email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		errMessage := constants.ErrUserNotFound
		fmt.Println("[GetUserByEmailPath]", errMessage)
		return models.User{}, errors.New(errMessage)
	} else if err != nil {
		fmt.Println("[GetUserByEmailPath]", err.Error())
		return models.User{}, err
	}

	fmt.Println("[GetUserByEmailPath] Returned user details from repository")

	return user, nil
}

func (us *UserService) Login(request *requests.Login) (*models.User, error) {
	fmt.Println("[LoginService] Hitting login function in user service")

	user, err := us.ur.GetUserByEmail(request.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessage := constants.ErrUserNotFound
		fmt.Println("[LoginService]", errorMessage)
		return nil, errors.New(errorMessage)
	}

	isValidPassword := utils.ComparePassword(user.Password, request.Password)

	if !isValidPassword {
		errMessage := constants.ErrWrongPassword
		fmt.Println("[LoginService]", errMessage)
		return nil, errors.New(errMessage)
	}

	fmt.Println("[LoginService] User login successful")
	return &user, nil
}

func (us *UserService) Signup(request *requests.Signup) (*models.User, error) {
	fmt.Println("[SignupService] Hitting signup function in user service")

	_, err := us.ur.GetUserByEmail(request.Email)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		const errorMessage = constants.ErrUserAlreadyExists
		fmt.Println("[SignupService]", errorMessage)
		return nil, errors.New(errorMessage)
	}

	hashedPassword, errHashedPassword := utils.GenerateHashPassword(request.Password)

	if errHashedPassword != nil {
		fmt.Println("[SignupService]", err.Error())
		return &models.User{}, errHashedPassword
	}

	user := &models.User{
		Email:     request.Email,
		Role:      constants.User,
		Password:  hashedPassword,
		LastName:  request.LastName,
		FirstName: request.FirstName,
	}

	createdUser, errCreatedUser := us.ur.CreateUser(user)

	if errCreatedUser != nil {
		fmt.Println("[SaveService]", errCreatedUser.Error())
		return &models.User{}, errCreatedUser
	}

	fmt.Println("[SignupService] Returned saved user details from repository")

	return &createdUser, nil
}

func (us *UserService) Update(email string, request *requests.Update) (models.User, error) {
	fmt.Println("[UpdateService] update user details in user service")

	user, err := us.ur.UpdateUserByEmail(email, request)

	if err != nil {
		fmt.Println("[UpdateService]", err.Error())
		return models.User{}, err
	}

	fmt.Println("[UpdateService], update user details in user service")

	return user, nil
}
