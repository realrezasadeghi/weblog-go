package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"weblog/models"
)

func LoadDatabaseConfig() (models.Database, error) {
	fmt.Println("[LoadDatabaseConfig] Loading env database variables")

	err := godotenv.Load()

	if err != nil {
		fmt.Println("[GetDSNString]", err.Error())
		return models.Database{}, err
	}

	config := models.Database{
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}

	return config, nil
}

func LoadSmtpConfigs() (models.Smtp, error) {
	fmt.Println("[LoadConfig] Loading env smtp variables")

	err := godotenv.Load()

	if err != nil {
		fmt.Println("[LoadSmtpConfigs]", err.Error())
		return models.Smtp{}, err
	}

	smtp := models.Smtp{
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPUser:     os.Getenv("SMTP_USER"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		EmailFrom:    os.Getenv("EMAIL_FROM"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}

	return smtp, nil
}

func GetDsn(config models.Database) string {
	fmt.Println("[GetDSNString] Getting dns string")
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	fmt.Println("[GetDSNString] Returning dsn string value")
	return dsn
}
