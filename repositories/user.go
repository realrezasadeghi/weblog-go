package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"weblog/models"
)

type IUserRepository interface {
	GetAllUsers() ([]models.User, error)
	DeleteUserByEmail(email string) error
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user *models.User) (models.User, error)
	UpdateUserByEmail(email string, user models.User) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
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

func (ur *UserRepository) DeleteUserByEmail(email string) error {
	fmt.Println("[DeleteUserByEmail] delete user by email")

	response := ur.db.Where("email = ?", email).Delete(&models.User{})

	if response.Error != nil {
		fmt.Println("[DeleteUserByEmail]", response.Error.Error())
		return response.Error
	}

	fmt.Println("[DeleteUserByEmail] delete user with email", email)

	return nil
}

func (ur *UserRepository) GetUserByEmail(email string) (models.User, error) {
	fmt.Println("[GetUserByEmail] find user details by username in user repository")
	var user models.User

	response := ur.db.Where("email = ?", email).First(&user)

	if response.Error != nil {
		fmt.Println("[GetUserByEmail]", response.Error.Error())
		return models.User{}, response.Error
	} else if response.RowsAffected == 0 {
		fmt.Println("[FindUserByUsernameRepository] User is not found with username")
		return models.User{}, gorm.ErrRecordNotFound
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

func (ur *UserRepository) UpdateUserByEmail(email string, user models.User) (models.User, error) {
	fmt.Println("[UpdateUserByEmail] Hitting update user by email in user repository")

	response := ur.db.Model(&user).Updates(user)

	if response.Error != nil {
		fmt.Println("[UpdateUserByEmail]", response.Error.Error())
		return models.User{}, response.Error
	}

	fmt.Println("[UpdateUserByEmail] Updating user successful")

	return user, nil
}
