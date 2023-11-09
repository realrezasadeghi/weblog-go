package responses

type UserResponse struct {
	ID        uint   `json:"id"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}
