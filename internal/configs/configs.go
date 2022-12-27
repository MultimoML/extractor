package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() string {
	environment := os.Getenv("ENVIRONMENT")

	switch environment {
	case "dev":
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			fmt.Println("Loaded 'dev' env.")
		}
	case "prod":
		fmt.Println("Loaded 'prod' env.")
	default:
		log.Fatal("Env variable ENVIRONMENT is not set or set to wrong value")
	}

	return environment
}
