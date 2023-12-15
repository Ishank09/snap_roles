package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigInterface interface {
	LoadConfig() (*AppConfig, error)
}
type AppConfig struct {
	DB_SERVER   string
	DB_PORT     int64
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func LoadConfig() (*AppConfig, error) {
	var targetEnvironment string = os.Getenv("APP_ENV")
	var envFile string = ""
	switch targetEnvironment {
	case "dev":
		envFile = "dev.env"
	case "prod":
		envFile = "prod.env"
	default:
		envFile = "dev.env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		return nil, err
	}
	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	var config *AppConfig = &AppConfig{
		DB_SERVER:   os.Getenv("DB_SERVER"),
		DB_PORT:     dbPort,
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
	return config, nil

}
