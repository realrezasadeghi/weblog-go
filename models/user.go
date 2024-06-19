package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model         `json:"gorm.Model"`
	Role               string    `json:"role"`
	Photo              string    `json:"photo"`
	IsVerified         bool      `json:"verified"`
	Password           string    `json:"password"`
	LastName           string    `json:"last_name"`
	Email              string    `gorm:"primaryKey"`
	FirstName          string    `json:"first_name"`
	VerificationCode   string    `json:"verification_code"`
	PasswordResetAt    time.Time `json:"password_reset_at"`
	PasswordResetToken string    `json:"password_reset_token"`
}
