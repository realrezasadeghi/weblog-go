package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"weblog/models"
	"weblog/requests"
)

type IUserRepository interface {
	DeleteUserById(id string) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (models.User, error)
	UpdateUserById(id string, user *requests.Update) (models.User, error)
	UpdateUserResetPassword(email string, password *models.UpdatePassword) error
}

type UserRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) DeleteUserById(id string) error {
	fmt.Println("[DeleteUserById] delete user by id in user repository")

	response := ur.db.Where("id = ?", id).Delete(&models.User{})

	if response.Error != nil {
		fmt.Println("[DeleteUserById]", response.Error.Error())
		return response.Error
	}

	fmt.Println("[DeleteUserById] delete user with id in user repository", id)

	return nil
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	fmt.Println("[GetAllUsersController] find all users without condition in user repository")
	var users []models.User
	response := ur.db.Find(&users)

	if response.Error != nil {
		fmt.Println("[GetAllUsersController]", response.Error.Error())
		return []models.User{}, response.Error
	}

	fmt.Println("[GetAllUsersController] List of users has found in user repository")
	return users, nil
}

func (ur *UserRepository) GetUserById(id string) (models.User, error) {
	fmt.Println("[GetUserById] Hitting get user by id in user repository")

	var user models.User

	response := ur.db.Where("id = ?", id).Find(&user)

	if response.Error != nil {
		fmt.Println("[GetUserById]", response.Error.Error())
		return models.User{}, response.Error
	} else if response.RowsAffected == 0 {
		fmt.Println("[GetUserById] User is not found with id")
		return models.User{}, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	fmt.Println("[GetUserByEmail] find user details by username in user repository")

	var user *models.User

	response := ur.db.Where("email = ?", email).First(&user)

	if response.Error != nil {
		fmt.Println("[GetUserByEmail]", response.Error.Error())
		return &models.User{}, response.Error
	} else if response.RowsAffected == 0 {
		fmt.Println("[FindUserByUsernameRepository] User is not found with username")
		return &models.User{}, gorm.ErrRecordNotFound
	}

	fmt.Println("[GetUserByEmail] User detail has found")

	return user, nil
}

func (ur *UserRepository) CreateUser(user *models.User) (models.User, error) {
	fmt.Println("[CreateUser] Hitting save function in user repository")

	response := ur.db.Create(&user)

	if response.Error != nil {
		fmt.Println("[CreateUser]", response.Error.Error())
		return models.User{}, response.Error
	}

	fmt.Println("[CreateUser] Create user successful")

	return *user, nil
}

func (ur *UserRepository) UpdateUserById(id string, user *requests.Update) (models.User, error) {
	fmt.Println("[UpdateUserById] Hitting update user by email in user repository")

	var userResponse models.User
	response := ur.db.Where("id = ?", id).Model(&models.User{}).Updates(&user)

	if response.Error != nil {
		fmt.Println("[UpdateUserById]", response.Error.Error())
		return userResponse, response.Error
	}

	fmt.Println("[UpdateUserById] Updating user successful")

	response.Where("id = ?", id).Find(&userResponse)

	return userResponse, nil
}

func (ur *UserRepository) UpdateUserResetPassword(email string, password *models.UpdatePassword) error {

	response := ur.db.Where("email = ?", email).Model(&models.User{}).Updates(&password)

	if response.Error != nil {
		fmt.Println("[UpdateUserResetPassword]", response.Error.Error())
		return response.Error
	}

	return nil
}
