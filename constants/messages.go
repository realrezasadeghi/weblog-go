package constants

const (
	ErrInvalidId              = "invalid id"
	ErrInvalidToken           = "invalid token"
	ErrInvalidEmail           = "invalid email"
	ErrWrongPassword          = "password is wrong"
	ErrExpiredToken           = "token has expired"
	ErrTokensNotFound         = "tokens are not found"
	GetUserByEmailSuccessful  = "user received successful"
	DeleteUserByIdSuccessful  = "User deleted by id successful"
	LoginSuccessful           = "Your login was successful"
	SignupSuccessful          = "Your signup was successful"
	ErrMaliciousToken         = "malicious token has been passed"
	GetAllUsersSuccessful     = "User list received successfully"
	ErrUserNotFound           = "user is not found with username"
	UpdateSuccessful          = "You update information successful"
	ErrNoAuthHeader           = "no authorization header provided"
	ErrUserIsNotAuthorized    = "user is not authorized to this api"
	ErrUserAlreadyExists      = "user is already exists with email"
	ErrInvalidTokenExpiration = "invalid jwt access token expiration in minutes value"
)
