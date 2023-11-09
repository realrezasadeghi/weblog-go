package database

import (
	"fmt"
	"weblog/configs"
	"weblog/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config models.Database) *gorm.DB {
	fmt.Println("[InitDatabase] Initiating database")
	dsn := configs.GetDsn(config)
	fmt.Println("[InitDatabase] Opening the database connection")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("[InitDatabase]", err.Error())
	} else if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("[InitDatabase]", err.Error())
	}

	fmt.Println("[InitDatabase] Database connection has established")

	return db
}
