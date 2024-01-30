package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	envPath := "./../.env"
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
