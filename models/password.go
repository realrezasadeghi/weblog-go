package models

import "time"

type UpdatePassword struct {
	Password           string    `json:"password"`
	PasswordResetAt    time.Time `json:"password_reset_at"`
	PasswordResetToken string    `json:"password_reset_token"`
}
