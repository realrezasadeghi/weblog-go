package configs

import (
	"fmt"
	"os"
	"weblog/models"

	"github.com/joho/godotenv"
)

func LoadConfigs() (models.Database, error) {
	fmt.Println("[LoadConfig] Loading env variables")

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

func GetDsn(config models.Database) string {
	fmt.Println("[GetDSNString] Getting dns string")
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	fmt.Println("[GetDSNString] Returning dsn string value")
	return dsn
}
