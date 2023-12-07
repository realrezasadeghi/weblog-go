package requests

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Signup struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

type Update struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

type ForgetPassword struct {
	Email string `json:"email" binding:"required"`
}

type ResetPassword struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}
