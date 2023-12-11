package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_HOST string
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
		log.Fatal("Environment file not found, please check APP_ENV environment variable and specify the environemnt file path.")
		return nil, err
	}
	var config *AppConfig = &AppConfig{
		DB_HOST: os.Getenv("DB_HOST"),
	}
	return config, nil

}
