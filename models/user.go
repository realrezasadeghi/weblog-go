package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"gorm.Model"`
	Role       string `json:"role"`
	Email      string `gorm:"primaryKey"`
	Password   string `json:"password"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
}
